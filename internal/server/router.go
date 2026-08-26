package server

import (
	"net/http"
	
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/handler"
)

type Handlers struct {
	Link *handler.LinkHandler
}

func NewRouter(h Handlers) http.Handler {
	mux := http.NewServeMux()

	registerLinkRouter(mux, h.Link)

	return mux
}
