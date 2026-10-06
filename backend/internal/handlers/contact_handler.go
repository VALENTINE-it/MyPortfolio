package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"
)

type ContactHandler struct {
	service *services.ContactService
}

func NewContactHandler(service *services.ContactService) *ContactHandler {
	return &ContactHandler{
		service: service,
	}
}

// HandleContact handles POST /api/contact with validation and honeypot protection.
func (h *ContactHandler) HandleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		WriteError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
		return
	}

	// 1. Content-Type check
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(strings.ToLower(contentType), "application/json") {
		WriteError(w, http.StatusUnsupportedMediaType, "UNSUPPORTED_MEDIA_TYPE", "Content-Type must be application/json")
		return
	}

	// 2. Decode and validate single JSON object without unknown fields
	var req models.ContactRequest
	if err := DecodeJSON(w, r, &req); err != nil {
		if strings.Contains(err.Error(), "payload too large") {
			WriteError(w, http.StatusRequestEntityTooLarge, "PAYLOAD_TOO_LARGE", "Request payload exceeds maximum allowed size (64KB)")
			return
		}
		WriteError(w, http.StatusBadRequest, "INVALID_JSON", err.Error())
		return
	}

	// 3. Submit contact message with validation and honeypot check
	saved, err := h.service.SubmitContact(r.Context(), &req)
	if err != nil {
		var valErrs services.ValidationErrors
		if errors.As(err, &valErrs) {
			WriteError(w, http.StatusBadRequest, "VALIDATION_FAILED", "Validation failed", valErrs)
			return
		}

		slog.Error("Internal error saving contact message", "error", err)
		WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Unable to process your request at this time. Please try again later.")
		return
	}

	WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Thank you for reaching out! Your message has been received.",
		"data": map[string]interface{}{
			"id":        saved.ID,
			"createdAt": saved.CreatedAt,
		},
	})
}
