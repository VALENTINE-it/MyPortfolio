package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"portfolio-backend/internal/cache"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
	"portfolio-backend/internal/utils"
)

type ProjectHandler struct {
	service *services.ProjectService
	cache   cache.Cache
}

func NewProjectHandler(service *services.ProjectService, c cache.Cache) *ProjectHandler {
	return &ProjectHandler{
		service: service,
		cache:   c,
	}
}

// HandleProjects handles public GET /api/projects and GET /api/projects/{id}
func (h *ProjectHandler) HandleProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/projects")
	path = strings.Trim(path, "/")

	if path != "" {
		id, err := strconv.ParseInt(path, 10, 64)
		if err != nil || id <= 0 {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid project ID: must be a positive integer")
			return
		}

		project, err := h.service.GetProjectByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, repositories.ErrProjectNotFound) {
				utils.SendError(w, http.StatusNotFound, utils.CodeNotFound, "Project not found")
				return
			}
			slog.Error("Error fetching project by ID", "id", id, "error", err)
			utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Unable to retrieve project details")
			return
		}

		utils.SendSuccess(w, http.StatusOK, project)
		return
	}

	// GET all projects
	cacheKey := "projects:all"
	if h.cache != nil {
		if cached, found := h.cache.Get(r.Context(), cacheKey); found {
			etag := generateETag(cached)
			if r.Header.Get("If-None-Match") == etag {
				w.WriteHeader(http.StatusNotModified)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("ETag", etag)
			w.Header().Set("Cache-Control", "public, max-age=300")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(cached)
			return
		}
	}

	projects, err := h.service.GetAllProjects(r.Context())
	if err != nil {
		slog.Error("Error fetching all projects", "error", err)
		utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Unable to retrieve projects list")
		return
	}

	respBytes, err := json.Marshal(utils.SuccessResponse{
		Success: true,
		Data:    projects,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Failed to serialize response")
		return
	}

	if h.cache != nil {
		h.cache.Set(r.Context(), cacheKey, respBytes, 10*time.Minute)
	}

	etag := generateETag(respBytes)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("ETag", etag)
	w.Header().Set("Cache-Control", "public, max-age=300")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBytes)
}

// HandleAdminProjects handles protected admin CRUD routes under /api/admin/projects
func (h *ProjectHandler) HandleAdminProjects(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/admin/projects")
	path = strings.Trim(path, "/")

	switch r.Method {
	case http.MethodPost:
		if path != "" {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid endpoint for project creation")
			return
		}
		var input models.ProjectInput
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&input); err != nil {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, fmt.Sprintf("Invalid JSON body: %v", err))
			return
		}

		created, err := h.service.CreateProject(r.Context(), &input)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, err.Error())
			return
		}

		// Invalidate cache
		if h.cache != nil {
			h.cache.Delete(r.Context(), "projects:all")
		}

		utils.SendSuccessWithMessage(w, http.StatusCreated, "Project created successfully", created)

	case http.MethodPut:
		if path == "" {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Project ID is required in URL path")
			return
		}
		id, err := strconv.ParseInt(path, 10, 64)
		if err != nil || id <= 0 {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid project ID")
			return
		}

		var input models.ProjectInput
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&input); err != nil {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, fmt.Sprintf("Invalid JSON body: %v", err))
			return
		}

		updated, err := h.service.UpdateProject(r.Context(), id, &input)
		if err != nil {
			if errors.Is(err, repositories.ErrProjectNotFound) {
				utils.SendError(w, http.StatusNotFound, utils.CodeNotFound, "Project not found")
				return
			}
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, err.Error())
			return
		}

		// Invalidate cache
		if h.cache != nil {
			h.cache.Delete(r.Context(), "projects:all")
		}

		utils.SendSuccessWithMessage(w, http.StatusOK, "Project updated successfully", updated)

	case http.MethodDelete:
		if path == "" {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Project ID is required in URL path")
			return
		}
		id, err := strconv.ParseInt(path, 10, 64)
		if err != nil || id <= 0 {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid project ID")
			return
		}

		if err := h.service.DeleteProject(r.Context(), id); err != nil {
			if errors.Is(err, repositories.ErrProjectNotFound) {
				utils.SendError(w, http.StatusNotFound, utils.CodeNotFound, "Project not found")
				return
			}
			utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Failed to delete project")
			return
		}

		// Invalidate cache
		if h.cache != nil {
			h.cache.Delete(r.Context(), "projects:all")
		}

		utils.SendSuccessWithMessage(w, http.StatusOK, "Project deleted successfully", map[string]int64{"id": id})

	default:
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
	}
}

func generateETag(data []byte) string {
	h := sha256.Sum256(data)
	return fmt.Sprintf(`"%s"`, hex.EncodeToString(h[:8]))
}
