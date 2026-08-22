package app

import (
	"log"

	"google.golang.org/grpc"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/client"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/config"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/server"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/service"
)


func Run() {
	// cfg := config.Load()

	// linkConn, err := grpc.NewClient(cfg.LinkGRPC)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// linkClient := client.NewLinkClient(linkConn)

	// linkService := service.NewLinkService(linkClient)

	// linkHandler := handler.NewLinkhandler(*linkService)

	// router := server.NewRouter()

	// httpServer := server.New()
}