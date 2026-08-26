package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/client"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/config"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/server"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/service"
)

func New(cfg *config.Config) (server.Server, error) {
	linkConn, err := grpc.NewClient(cfg.LinkGRPC)
	if err != nil {
		log.Fatal(err)
	}

	linkClient := client.NewLinkClient(linkConn)

	linkService := service.NewLinkService(linkClient)

	linkHandler := handler.NewLinkhandler(*linkService)

	router := server.NewRouter(server.Handlers{
		Link: linkHandler,
	})

	httpServer := server.New(router, *cfg)

	return *httpServer, nil
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
