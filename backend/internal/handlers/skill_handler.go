package handlers

import (
	"log/slog"
	"net/http"

	"portfolio-backend/internal/services"
)

type SkillHandler struct {
	service *services.SkillService
}

func NewSkillHandler(service *services.SkillService) *SkillHandler {
	return &SkillHandler{service: service}
}

// HandleSkills handles GET /api/skills
func (h *SkillHandler) HandleSkills(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		WriteError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
		return
	}

	category := r.URL.Query().Get("category")
	groupedQuery := r.URL.Query().Get("grouped")

	if category != "" {
		skills, err := h.service.GetAllSkills(r.Context(), category)
		if err != nil {
			slog.Error("Error fetching skills by category", "category", category, "error", err)
			WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Unable to retrieve skills")
			return
		}

		WriteJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"data":    skills,
		})
		return
	}

	if groupedQuery == "false" {
		raw, err := h.service.GetAllSkills(r.Context(), "")
		if err != nil {
			slog.Error("Error fetching raw skills", "error", err)
			WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Unable to retrieve skills")
			return
		}

		WriteJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"data":    raw,
		})
		return
	}

	// Default: return grouped categories
	grouped, err := h.service.GetSkillsGrouped(r.Context())
	if err != nil {
		slog.Error("Error fetching grouped skills", "error", err)
		WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Unable to retrieve skills")
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    grouped,
	})
}
