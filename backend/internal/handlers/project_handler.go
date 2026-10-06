package handlers

import (
	"errors"
	"log/slog"
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
		w.Header().Set("Allow", http.MethodGet)
		WriteError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
		return
	}

	// Check if a specific ID was passed in URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/projects")
	path = strings.Trim(path, "/")

	if path != "" {
		id, err := strconv.ParseInt(path, 10, 64)
		if err != nil || id <= 0 {
			WriteError(w, http.StatusBadRequest, "INVALID_PROJECT_ID", "Invalid project ID: must be a positive integer")
			return
		}

		project, err := h.service.GetProjectByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, repositories.ErrProjectNotFound) {
				WriteError(w, http.StatusNotFound, "PROJECT_NOT_FOUND", "Project not found")
				return
			}

			slog.Error("Error fetching project by ID", "id", id, "error", err)
			WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Unable to retrieve project details")
			return
		}

		WriteJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"data":    project,
		})
		return
	}

	// GetAllProjects
	projects, err := h.service.GetAllProjects(r.Context())
	if err != nil {
		slog.Error("Error fetching projects", "error", err)
		WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Unable to retrieve projects list")
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    projects,
	})
}
