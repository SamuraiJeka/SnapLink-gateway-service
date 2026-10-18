package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

type App struct {
	server *server.Server
	linkConn *grpc.ClientConn
	authConn *grpc.ClientConn
}

func (a *App) Start() error {
	return a.server.Start()
}

func (a *App) Shutdown(ctx context.Context) error {
	httpErr := a.server.Shutdown(ctx)

	authErr := a.authConn.Close()
	linkErr := a.linkConn.Close()

	return errors.Join(httpErr, authErr, linkErr)
}

func New(cfg *config.Config) (*App, error) {
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
		_ = authConn.Close()
		_ = linkConn.Close()
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

	return &App{
		server: httpServer,
		linkConn: linkConn,
		authConn: authConn,
	}, nil
}

func Run() {
	cfg := config.Load()

	application, err := New(cfg)
	if err != nil {
		log.Printf("initialize application: %v", err)
		return
	}

	serverErr := make(chan error, 1)

	go func() {
		log.Printf("Server started on %s", cfg.HttpAddr)
		serverErr <- application.Start()
	}()

	signalCtx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	select {
	case err := <-serverErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server error: %v", err)
		}
	case <- signalCtx.Done():
		log.Println("Shutdown signal received")

		ctx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := application.Shutdown(ctx); err != nil {
			log.Printf("Shutdown error: %v", err)
		}
	}

	log.Println("Application stopped")
}
