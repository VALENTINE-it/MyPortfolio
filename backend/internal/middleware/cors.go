package middleware

import (
	"net/http"
	"strings"
)

// CORS handles Cross-Origin Resource Sharing with strict origin verification and credentials support
func CORS(allowedOrigins []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin != "" {
			isAllowed := false
			for _, allowed := range allowedOrigins {
				if allowed == "*" || allowed == origin || strings.TrimRight(allowed, "/") == strings.TrimRight(origin, "/") {
					isAllowed = true
					break
				}
			}

			if isAllowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID, X-CSRF-Token")
				w.Header().Set("Access-Control-Expose-Headers", "X-Request-ID, ETag")
				w.Header().Set("Access-Control-Max-Age", "86400")
			} else if r.Method == http.MethodOptions {
				// Reject unauthorized preflight
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
