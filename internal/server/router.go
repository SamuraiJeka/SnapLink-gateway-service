package server

import (
	"net/http"
	
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
)


func NewRouter(
	linkHandler *handler.LinkHandler,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /api/v1/links",
		linkHandler.Create,
	)

	var handler http.Handler = mux

	return handler
}