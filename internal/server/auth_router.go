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
		"POST /auth/login",
		h.Login,
	)

	mux.HandleFunc(
		"POST /auth/refresh",
		h.Refresh,
	)

	mux.HandleFunc(
		"POST /auth/logout",
		h.Logout,
	)

	mux.HandleFunc(
		"POST /auth/validate",
		h.ValidateAccessToken,
	)
}
