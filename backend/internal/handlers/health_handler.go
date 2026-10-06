package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// HandleHealth handles GET /api/health returning minimal status.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		WriteError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(HealthResponse{
		Status: "ok",
	})
}
