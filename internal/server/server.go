package server

import (
	"context"
	"net/http"
	"time"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/config"
)

type Server struct {
	httpServer * http.Server
}

func New(
	handler http.Handler,
	cfg config.Config,
) *Server {

	return &Server{
		httpServer: &http.Server{
			Addr: cfg.HttpAddr,
			Handler: handler,
			ReadTimeout: 5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout: 60 * time.Second,
		},
	}
}


func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}


func (s *Server) Shutdown(
	ctx context.Context,
) error {
	return s.httpServer.Shutdown(ctx)
}
