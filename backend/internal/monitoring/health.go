package monitoring

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"portfolio-backend/internal/utils"
)

// HealthHandler provides liveness and readiness probe endpoints
type HealthHandler struct {
	db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

// LivenessHandler handles GET /health/live (Process is running)
func (h *HealthHandler) LivenessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"status":    "alive",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// ReadinessHandler handles GET /health/ready (Checks database connectivity)
func (h *HealthHandler) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if h.db == nil {
		utils.SendError(w, http.StatusServiceUnavailable, utils.CodeServiceUnavailable, "Database connection not initialized")
		return
	}

	start := time.Now()
	if err := h.db.PingContext(ctx); err != nil {
		GetMetrics().RecordDBQuery(time.Since(start), err)
		utils.WriteJSON(w, http.StatusServiceUnavailable, map[string]interface{}{
			"status":    "unhealthy",
			"database":  "down",
			"error":     err.Error(),
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	GetMetrics().RecordDBQuery(time.Since(start), nil)

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"status":    "ready",
		"database":  "up",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// LegacyHealthHandler handles GET /api/health for frontend backwards compatibility
func (h *HealthHandler) LegacyHealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "API is running",
	})
}
