package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
	"portfolio-backend/internal/utils"
)

type ContactHandler struct {
	service *services.ContactService
}

func NewContactHandler(service *services.ContactService) *ContactHandler {
	return &ContactHandler{
		service: service,
	}
}

// HandleContact handles public POST /api/contact
func (h *ContactHandler) HandleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}

	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(strings.ToLower(contentType), "application/json") {
		utils.SendError(w, http.StatusUnsupportedMediaType, utils.CodeUnsupportedMedia, "Content-Type must be application/json")
		return
	}

	var req models.ContactRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Malformed JSON payload or invalid request body format")
		return
	}

	saved, err := h.service.SubmitContact(r.Context(), &req)
	if err != nil {
		var valErrs services.ValidationErrors
		if errors.As(err, &valErrs) {
			utils.SendErrorWithDetails(w, http.StatusBadRequest, utils.CodeValidationError, "Validation failed", valErrs)
			return
		}

		slog.Error("Internal error submitting contact message", "error", err)
		utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Unable to process your request at this time. Please try again later.")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Thank you for reaching out! Your message has been received.",
		"data": map[string]interface{}{
			"id":        saved.ID,
			"createdAt": saved.CreatedAt,
		},
	})
}

// HandleAdminMessages handles admin routes under /api/admin/messages
func (h *ContactHandler) HandleAdminMessages(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/admin/messages")
	path = strings.Trim(path, "/")

	switch r.Method {
	case http.MethodGet:
		if path != "" {
			// GET /api/admin/messages/{id}
			id, err := strconv.ParseInt(path, 10, 64)
			if err != nil || id <= 0 {
				utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid message ID")
				return
			}
			msg, err := h.service.GetMessageByID(r.Context(), id)
			if err != nil {
				if errors.Is(err, repositories.ErrContactNotFound) {
					utils.SendError(w, http.StatusNotFound, utils.CodeNotFound, "Message not found")
					return
				}
				utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Failed to retrieve message")
				return
			}
			utils.SendSuccess(w, http.StatusOK, msg)
			return
		}

		// GET /api/admin/messages (paginated)
		limit := 20
		offset := 0
		if lStr := r.URL.Query().Get("limit"); lStr != "" {
			if l, err := strconv.Atoi(lStr); err == nil && l > 0 {
				limit = l
			}
		}
		if oStr := r.URL.Query().Get("offset"); oStr != "" {
			if o, err := strconv.Atoi(oStr); err == nil && o >= 0 {
				offset = o
			}
		}
		status := r.URL.Query().Get("status")

		messages, total, err := h.service.GetMessages(r.Context(), limit, offset, status)
		if err != nil {
			slog.Error("Failed to fetch messages for admin", "error", err)
			utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Failed to retrieve messages")
			return
		}

		utils.SendSuccessWithMeta(w, http.StatusOK, messages, map[string]interface{}{
			"total":  total,
			"limit":  limit,
			"offset": offset,
		})

	case http.MethodPut:
		// PUT /api/admin/messages/{id}/status
		parts := strings.Split(path, "/")
		if len(parts) == 0 || parts[0] == "" {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Message ID is required")
			return
		}
		id, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil || id <= 0 {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid message ID")
			return
		}

		var payload struct {
			Status string `json:"status"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid JSON payload")
			return
		}

		if err := h.service.UpdateMessageStatus(r.Context(), id, payload.Status); err != nil {
			if errors.Is(err, repositories.ErrContactNotFound) {
				utils.SendError(w, http.StatusNotFound, utils.CodeNotFound, "Message not found")
				return
			}
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, err.Error())
			return
		}

		utils.SendSuccessWithMessage(w, http.StatusOK, "Message status updated successfully", map[string]interface{}{
			"id":     id,
			"status": payload.Status,
		})

	case http.MethodDelete:
		// DELETE /api/admin/messages/{id}
		id, err := strconv.ParseInt(path, 10, 64)
		if err != nil || id <= 0 {
			utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid message ID")
			return
		}

		if err := h.service.DeleteMessage(r.Context(), id); err != nil {
			if errors.Is(err, repositories.ErrContactNotFound) {
				utils.SendError(w, http.StatusNotFound, utils.CodeNotFound, "Message not found")
				return
			}
			utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "Failed to delete message")
			return
		}

		utils.SendSuccessWithMessage(w, http.StatusOK, "Message deleted successfully", map[string]int64{"id": id})

	default:
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
	}
}
