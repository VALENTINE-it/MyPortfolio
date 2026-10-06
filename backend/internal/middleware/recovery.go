package middleware

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
)

// PanicRecovery catches unhandled panics, logs the stack trace server-side via slog,
// and returns a standard HTTP 500 JSON response without crashing the server.
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				reqID := GetRequestID(r.Context())
				stack := string(debug.Stack())

				slog.Error("CRITICAL: Panic recovered in HTTP handler",
					"request_id", reqID,
					"error", fmt.Sprintf("%v", rec),
					"stack", stack,
					"path", r.URL.Path,
				)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"error": map[string]interface{}{
						"code":    "INTERNAL_SERVER_ERROR",
						"message": "An internal server error occurred.",
					},
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
