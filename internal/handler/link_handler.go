package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/service"
)


type LinkHandler struct {
	service service.LinkService
}

func NewLinkhandler(
	service service.LinkService,
) *LinkHandler {
	return &LinkHandler{
		service: service,
	}
}

func (h *LinkHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateLink

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	link, err := h.service.Create(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(link); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
		return
	}
}


func (h *LinkHandler) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	link, err := h.service.Get(r.Context(), idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(link); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
		return
	}
}

func (h *LinkHandler) GetByUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query()

	limit := 20
	offset := 0

	if value := query.Get("limit"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed <= 0 {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}
		limit = parsed
	}

	if value := query.Get("offset"); value != "" {
		parsed , err := strconv.Atoi(value)
		if err != nil || parsed < 0 {
			http.Error(w, "invalid offset", http.StatusBadRequest)
			return
		}
		offset = parsed
	}

	links, err := h.service.GetByUser(r.Context(), uint32(limit), uint32(offset))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(links); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
		return
	}
}

func (h *LinkHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(r.Context(), idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode("ok")
}
