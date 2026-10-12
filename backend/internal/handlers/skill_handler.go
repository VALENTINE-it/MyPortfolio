package handlers

import (
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

type SkillHandler struct {
	service *services.SkillService
	cache   cache.Cache
}

func NewSkillHandler(service *services.SkillService, c cache.Cache) *SkillHandler {
	return &SkillHandler{
		service: service,
		cache:   c,
	}
}

// HandleSkills handles public GET /api/skills
func (h *SkillHandler) HandleSkills(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}

	category := r.URL.Query().Get("category")
	groupedQuery := r.URL.Query().Get("grouped")

	cacheKey := fmt.Sprintf("skills:cat=%s:grouped=%s", category, groupedQuery)

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

	var data interface{}
	var err error

	if category != "" {
		data, err = h.service.GetAllSkills(r.Context(), category)
	} else if groupedQuery == "false" {
		data, err = h.service.GetAllSkills(r.Context(), "")
	} else {
		data, err = h.service.GetSkillsGrouped(r.Context())
	}

	if err != nil {
		slog.Error("Error fetching skills", "category", category, "error", err)
		utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Unable to retrieve skills")
		return
	}

	respBytes, err := json.Marshal(utils.SuccessResponse{
		Success: true,
		Data:    data,
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

// HandleAdminSkills handles protected admin CRUD routes under /api/admin/skills
func (h *SkillHandler) HandleAdminSkills(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/admin/skills")
	path = strings.Trim(path, "/")

	switch r.Method {
	case http.MethodPost:
		if path != "" {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid endpoint for skill creation")
			return
		}

		var input models.SkillInput
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&input); err != nil {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, fmt.Sprintf("Invalid JSON body: %v", err))
			return
		}

		created, err := h.service.CreateSkill(r.Context(), &input)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, err.Error())
			return
		}

		if h.cache != nil {
			h.cache.DeletePrefix(r.Context(), "skills:")
		}

		utils.SendSuccessWithMessage(w, http.StatusCreated, "Skill created successfully", created)

	case http.MethodPut:
		if path == "" {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Skill ID is required in URL path")
			return
		}
		id, err := strconv.ParseInt(path, 10, 64)
		if err != nil || id <= 0 {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid skill ID")
			return
		}

		var input models.SkillInput
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&input); err != nil {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, fmt.Sprintf("Invalid JSON body: %v", err))
			return
		}

		updated, err := h.service.UpdateSkill(r.Context(), id, &input)
		if err != nil {
			if errors.Is(err, repositories.ErrSkillNotFound) {
				utils.SendError(w, http.StatusNotFound, utils.CodeNotFound, "Skill not found")
				return
			}
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, err.Error())
			return
		}

		if h.cache != nil {
			h.cache.DeletePrefix(r.Context(), "skills:")
		}

		utils.SendSuccessWithMessage(w, http.StatusOK, "Skill updated successfully", updated)

	case http.MethodDelete:
		if path == "" {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Skill ID is required in URL path")
			return
		}
		id, err := strconv.ParseInt(path, 10, 64)
		if err != nil || id <= 0 {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid skill ID")
			return
		}

		if err := h.service.DeleteSkill(r.Context(), id); err != nil {
			if errors.Is(err, repositories.ErrSkillNotFound) {
				utils.SendError(w, http.StatusNotFound, utils.CodeNotFound, "Skill not found")
				return
			}
			utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Failed to delete skill")
			return
		}

		if h.cache != nil {
			h.cache.DeletePrefix(r.Context(), "skills:")
		}

		utils.SendSuccessWithMessage(w, http.StatusOK, "Skill deleted successfully", map[string]int64{"id": id})

	default:
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
	}
}
