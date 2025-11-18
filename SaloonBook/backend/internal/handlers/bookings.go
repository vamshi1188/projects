package handlers

import (
	"encoding/json"
	"net/http"
	"saloonbook/internal/models"
	"saloonbook/internal/repository"
)

type BookingHandler struct {
	repo *repository.BookingRepository
}

func NewBookingHandler(repo *repository.BookingRepository) *BookingHandler {
	return &BookingHandler{repo: repo}
}

func (h *BookingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if booking.Customer == "" || booking.Phone == "" {
		http.Error(w, "customer and phone are required", http.StatusBadRequest)
		return
	}

	created, err := h.repo.Create(r.Context(), booking)
	if err != nil {
		http.Error(w, "failed to create booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *BookingHandler) List(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.repo.List(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch bookings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}
