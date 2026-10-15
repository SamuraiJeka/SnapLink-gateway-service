package server

import (
	"net/http"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/middleware"
)

type Handlers struct {
	Link *handler.LinkHandler
	Auth *handler.AuthHandler
	Middleware middleware.AuthMiddleware
}

func NewRouter(h Handlers) http.Handler {
	mux := http.NewServeMux()

	registerLinkRouter(mux, h.Link, h.Middleware)
	registerAuthRouter(mux, h.Auth)

	return mux
}
