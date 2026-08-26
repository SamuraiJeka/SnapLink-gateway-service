package server

import (
	"net/http"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
)

func registerLinkRouter(mux *http.ServeMux, h *handler.LinkHandler) {
	mux.HandleFunc(
		"POST /link/",
		h.Create,
	)

	mux.HandleFunc(
		"GET /link/{id}",
		h.Get,
	)

	mux.HandleFunc(
		"GET /link/user/{id}",
		h.GetByUser,
	)

	mux.HandleFunc(
		"DELETE /link/{id}",
		h.Delete,
	)
}
