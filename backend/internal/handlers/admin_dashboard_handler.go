package handlers

import (
	"database/sql"
	"net/http"
	"runtime"
	"time"

	"portfolio-backend/internal/utils"
)

type AdminDashboardHandler struct {
	db        *sql.DB
	startTime time.Time
}

func NewAdminDashboardHandler(db *sql.DB) *AdminDashboardHandler {
	return &AdminDashboardHandler{
		db:        db,
		startTime: time.Now(),
	}
}

// HandleDashboard handles GET /api/admin/dashboard
func (h *AdminDashboardHandler) HandleDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}

	var projectCount, skillCount, messageCount, unreadMessages int

	_ = h.db.QueryRowContext(r.Context(), "SELECT COUNT(*) FROM projects").Scan(&projectCount)
	_ = h.db.QueryRowContext(r.Context(), "SELECT COUNT(*) FROM skills").Scan(&skillCount)
	_ = h.db.QueryRowContext(r.Context(), "SELECT COUNT(*) FROM contacts").Scan(&messageCount)
	_ = h.db.QueryRowContext(r.Context(), "SELECT COUNT(*) FROM contacts WHERE status = 'unread'").Scan(&unreadMessages)

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	utils.SendSuccess(w, http.StatusOK, map[string]interface{}{
		"stats": map[string]interface{}{
			"projects":       projectCount,
			"skills":         skillCount,
			"totalMessages":  messageCount,
			"unreadMessages": unreadMessages,
		},
		"system": map[string]interface{}{
			"uptimeSeconds": time.Since(h.startTime).Seconds(),
			"goroutines":    runtime.NumGoroutine(),
			"allocMB":       float64(mem.Alloc) / (1024 * 1024),
			"sysMB":         float64(mem.Sys) / (1024 * 1024),
		},
	})
}
