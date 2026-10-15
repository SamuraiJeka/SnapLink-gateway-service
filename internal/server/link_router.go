package server

import (
	"net/http"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/middleware"
)

func registerLinkRouter(
	mux *http.ServeMux,
	h *handler.LinkHandler,
	authMiddleware middleware.AuthMiddleware,
) {
	mux.Handle(
		"POST /link/",
		authMiddleware.Handle(
			http.HandlerFunc(h.Create),
		),
	)

	mux.Handle(
		"GET /link/{id}",
		authMiddleware.Handle(
			http.HandlerFunc(h.Get),
		),
	)

	mux.Handle(
		"GET /link/user/{id}",
		authMiddleware.Handle(
			http.HandlerFunc(h.GetByUser),
		),
	)

	mux.Handle(
		"DELETE /link/{id}",
		authMiddleware.Handle(
			http.HandlerFunc(h.Delete),
		),
	)
}
