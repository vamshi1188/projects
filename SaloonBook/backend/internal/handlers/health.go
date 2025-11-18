package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"saloonbook/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
)

type HealthHandler struct {
	conn *pgx.Conn
}

func NewHealthHandler(conn *pgx.Conn) *HealthHandler {
	return &HealthHandler{conn: conn}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	response := models.HealthResponse{
		Status:  "ok",
		Version: "1.0.0",
	}

	if err := h.conn.Ping(ctx); err != nil {
		response.Status = "error"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
