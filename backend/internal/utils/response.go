package utils

import (
	"encoding/json"
	"net/http"
)

// StandardError represents structured error detail
type StandardError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// ErrorResponse represents the centralized error response format
type ErrorResponse struct {
	Success bool          `json:"success"`
	Error   StandardError `json:"error"`
}

// SuccessResponse represents the standardized success envelope
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// WriteJSON writes any payload as application/json with the given HTTP status
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// SendSuccess sends a success response
func SendSuccess(w http.ResponseWriter, status int, data interface{}) {
	WriteJSON(w, status, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

// SendSuccessWithMessage sends a success response with a human-readable confirmation message
func SendSuccessWithMessage(w http.ResponseWriter, status int, message string, data interface{}) {
	WriteJSON(w, status, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendSuccessWithMeta sends a paginated/meta-enriched success response
func SendSuccessWithMeta(w http.ResponseWriter, status int, data interface{}, meta interface{}) {
	WriteJSON(w, status, SuccessResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

// SendError sends a standardized JSON error
func SendError(w http.ResponseWriter, status int, code, message string) {
	WriteJSON(w, status, ErrorResponse{
		Success: false,
		Error: StandardError{
			Code:    code,
			Message: message,
		},
	})
}

// SendErrorWithDetails sends a standardized JSON error with field-level details
func SendErrorWithDetails(w http.ResponseWriter, status int, code, message string, details interface{}) {
	WriteJSON(w, status, ErrorResponse{
		Success: false,
		Error: StandardError{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}
