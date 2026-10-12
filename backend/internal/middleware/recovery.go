package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	"portfolio-backend/internal/monitoring"
	"portfolio-backend/internal/utils"
)

// Recovery recovers from unhandled panics, logs stack traces with request ID, and returns HTTP 500
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				reqID := monitoring.GetRequestID(r.Context())
				stack := string(debug.Stack())

				slog.Error("Unhandled HTTP panic recovered",
					"error", fmt.Sprintf("%v", rec),
					"request_id", reqID,
					"path", r.URL.Path,
					"method", r.Method,
					"stack", stack,
				)

				monitoring.GetMetrics().RecordRequest(r.Method, r.URL.Path, http.StatusInternalServerError, 0)

				utils.SendError(w, http.StatusInternalServerError, utils.CodeInternalServerError, "An internal server error occurred. Please try again later.")
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// PanicRecovery is an alias for Recovery
func PanicRecovery(next http.Handler) http.Handler {
	return Recovery(next)
}
