package handlers

import (
	"encoding/json"
	"net/http"
	"saloonbook/internal/repository"
)

type ServiceHandler struct {
	repo *repository.ServiceRepository
}

func NewServiceHandler(repo *repository.ServiceRepository) *ServiceHandler {
	return &ServiceHandler{repo: repo}
}

func (h *ServiceHandler) List(w http.ResponseWriter, r *http.Request) {
	services, err := h.repo.List(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch services", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}
