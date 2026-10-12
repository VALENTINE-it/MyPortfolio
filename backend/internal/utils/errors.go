package utils

import "errors"

// Common Error Codes
const (
	CodeBadRequest          = "BAD_REQUEST"
	CodeValidationError     = "VALIDATION_FAILED"
	CodeUnauthorized        = "UNAUTHORIZED"
	CodeForbidden           = "FORBIDDEN"
	CodeNotFound            = "NOT_FOUND"
	CodeMethodNotAllowed    = "METHOD_NOT_ALLOWED"
	CodeConflict            = "CONFLICT"
	CodeTooManyRequests     = "TOO_MANY_REQUESTS"
	CodeUnsupportedMedia    = "UNSUPPORTED_MEDIA_TYPE"
	CodePayloadTooLarge     = "PAYLOAD_TOO_LARGE"
	CodeInternalServerError = "INTERNAL_SERVER_ERROR"
	CodeServiceUnavailable  = "SERVICE_UNAVAILABLE"
	CodeCSRFInvalid         = "CSRF_TOKEN_INVALID"
)

// Common Domain Errors
var (
	ErrNotFound            = errors.New("requested resource not found")
	ErrAlreadyExists       = errors.New("resource already exists")
	ErrInvalidCredentials  = errors.New("invalid username or password")
	ErrSessionExpired      = errors.New("session has expired")
	ErrUnauthorized        = errors.New("authentication required")
	ErrForbidden           = errors.New("permission denied")
	ErrDatabase            = errors.New("database operation failed")
	ErrValidation          = errors.New("validation failed")
	ErrPayloadTooLarge     = errors.New("request payload exceeds maximum limit")
	ErrRateLimitExceeded   = errors.New("rate limit exceeded")
)
