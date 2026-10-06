package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// APIResponse represents the standard standardized JSON API envelope.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

// APIError represents structured error details.
type APIError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// WriteJSON encodes data as JSON with Content-Type header.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// WriteSuccess writes a standard successful JSON envelope.
func WriteSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	WriteJSON(w, status, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// WriteError writes a structured error envelope.
func WriteError(w http.ResponseWriter, status int, code string, message string, details ...interface{}) {
	var det interface{}
	if len(details) > 0 {
		det = details[0]
	}

	WriteJSON(w, status, APIResponse{
		Success: false,
		Message: message,
		Error: &APIError{
			Code:    code,
			Message: message,
			Details: det,
		},
	})
}

// DecodeJSON safely parses a single JSON object from request body,
// disallowing unknown fields and verifying no extra trailing data.
func DecodeJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("malformed JSON at position %d", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("malformed JSON payload")
		case errors.As(err, &unmarshalTypeError):
			return fmt.Errorf("invalid value for field %q", unmarshalTypeError.Field)
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("unknown field %s", fieldName)
		case errors.As(err, &maxBytesError):
			return errors.New("request payload too large")
		case errors.Is(err, io.EOF):
			return errors.New("request body cannot be empty")
		default:
			return fmt.Errorf("invalid request body: %w", err)
		}
	}

	// Verify that there is no second JSON object or garbage trailing the payload
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("request body must contain exactly one JSON object")
	}

	return nil
}
