package handlers

import (
	"log"
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
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	category := r.URL.Query().Get("category")
	groupedQuery := r.URL.Query().Get("grouped")

	if category != "" {
		skills, err := h.service.GetAllSkills(category)
		if err != nil {
			log.Printf("Error fetching skills for category %s: %v", category, err)
			writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"message": "Internal server error",
			})
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"data":    skills,
		})
		return
	}

	// If grouped is requested or default
	grouped, err := h.service.GetSkillsGrouped()
	if err != nil {
		log.Printf("Error fetching grouped skills: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": "Internal server error",
		})
		return
	}

	if groupedQuery == "false" {
		raw, err := h.service.GetAllSkills("")
		if err != nil {
			log.Printf("Error fetching all raw skills: %v", err)
			writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"message": "Internal server error",
			})
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"data":    raw,
		})
		return
	}

	// By default return grouped categories, with individual count
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    grouped,
	})
}
