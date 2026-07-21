package handler

import (
	"context"
	"encoding/json"
	"net/http"
)


type LinkService interface {
	Create(
		ctx context.Context,
		url string,
	) (string, error)
}

type LinkHandler struct {
	service LinkService
}


func NewLinkhandler(
	service LinkService,
) *LinkHandler {
	return &LinkHandler{
		service: service,
	}
}

func (h *LinkHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req CreateLinkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"Invalid json",
			http.StatusBadRequest,
		)
		return
	}

	shortURL, err := h.service.Create(
		r.Context(),
		req.URL,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	resp := CreateLinkResponse{
		ShortURL: shortURL,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(resp)
}