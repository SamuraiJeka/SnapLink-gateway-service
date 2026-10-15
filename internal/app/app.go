package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/client"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/config"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
	keyprovider "github.com/SamuraiJeka/SnapLink-gateway-service/internal/key_provider"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/middleware"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/server"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/service"
)

func New(cfg *config.Config) (*server.Server, error) {
	linkConn, err := grpc.NewClient(cfg.LinkGRPC)
	if err != nil {
		return nil, fmt.Errorf("create link gRPC client: %w", err)
	}
	authConn, err := grpc.NewClient(cfg.AuthGRPC)
	if err != nil {
		return nil, fmt.Errorf("create auth gRPC client: %w", err)
	}

	linkClient := client.NewLinkClient(linkConn)
	authClient := client.NewAuthClient(authConn)

	ctx, cancel := context.WithTimeout(
	context.Background(),
	5*time.Second,
	)
	defer cancel()

	keyProvider := keyprovider.NewKeyProvider(*authClient)
	if err := keyProvider.Refresh(ctx); err != nil {
		return nil, err
	}

	verifier := keyprovider.NewVerifier(keyProvider)

	authMiddleware := middleware.NewAuthMiddleware(verifier)

	linkService := service.NewLinkService(linkClient)
	authService := service.NewAuthService(authClient)

	linkHandler := handler.NewLinkhandler(*linkService)
	authHandler := handler.NewAuthHandler(*authService)

	router := server.NewRouter(server.Handlers{
		Link: linkHandler,
		Auth: authHandler,
		Middleware: *authMiddleware,
	})

	httpServer := server.New(router, *cfg)

	return httpServer, nil
}

func Run() {
	cfg := config.Load()

	application, err := New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Printf("Server started on %s", cfg.HttpAddr)

		if err := application.Start(); err != nil &&
		err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<- quit

	log.Println("Shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := application.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error %v", err)
	}
	
	log.Println("Server stopped")
}
