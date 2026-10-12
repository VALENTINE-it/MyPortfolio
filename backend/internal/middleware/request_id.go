package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"portfolio-backend/internal/monitoring"
)

// RequestID attaches or generates a unique X-Request-ID for every incoming HTTP request
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			bytes := make([]byte, 16)
			if _, err := rand.Read(bytes); err == nil {
				reqID = hex.EncodeToString(bytes)
			} else {
				reqID = "req-fallback"
			}
		}

		ctx := monitoring.WithRequestID(r.Context(), reqID)
		w.Header().Set("X-Request-ID", reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequestIDMiddleware is an alias for RequestID
func RequestIDMiddleware(next http.Handler) http.Handler {
	return RequestID(next)
}

// GetRequestID returns the request ID from context
func GetRequestID(ctx context.Context) string {
	return monitoring.GetRequestID(ctx)
}
