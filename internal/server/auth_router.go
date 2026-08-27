package server

import (
	"net/http"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
)

func registerAuthRouter(mux *http.ServeMux, h *handler.AuthHandler) {
	mux.HandleFunc(
		"POST /auth/register",
		h.Register,
	)

	mux.HandleFunc(
		"GET /auth/login",
		h.Login,
	)

	mux.HandleFunc(
		"GET /auth/refresh",
		h.Refresh,
	)

	mux.HandleFunc(
		"GET /auth/logout",
		h.Logout,
	)

	mux.HandleFunc(
		"GET /auth/validate",
		h.ValidateAccessToken,
	)
}
