package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
)

type ProjectHandler struct {
	service *services.ProjectService
}

func NewProjectHandler(service *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

// HandleProjects routes GET /api/projects and GET /api/projects/{id}
func (h *ProjectHandler) HandleProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	// Check if an ID was provided via path or PathValue
	path := strings.TrimPrefix(r.URL.Path, "/api/projects")
	path = strings.Trim(path, "/")

	if path != "" {
		id, err := strconv.ParseInt(path, 10, 64)
		if err != nil || id <= 0 {
			writeJSON(w, http.StatusBadRequest, map[string]interface{}{
				"success": false,
				"message": "Invalid project ID: must be a positive integer",
			})
			return
		}

		project, err := h.service.GetProjectByID(id)
		if err != nil {
			if errors.Is(err, repositories.ErrProjectNotFound) {
				writeJSON(w, http.StatusNotFound, map[string]interface{}{
					"success": false,
					"message": "Project not found",
				})
				return
			}
			log.Printf("Error fetching project %d: %v", id, err)
			writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"message": "Internal server error",
			})
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"data":    project,
		})
		return
	}

	// GetAllProjects
	projects, err := h.service.GetAllProjects()
	if err != nil {
		log.Printf("Error fetching all projects: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": "Internal server error",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    projects,
	})
}

// writeJSON is a helper to serialize response payload
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
