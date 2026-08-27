package server

import (
	"net/http"
	
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
)

type Handlers struct {
	Link *handler.LinkHandler
	Auth *handler.AuthHandler
}

func NewRouter(h Handlers) http.Handler {
	mux := http.NewServeMux()

	registerLinkRouter(mux, h.Link)
	registerAuthRouter(mux, h.Auth)

	return mux
}
