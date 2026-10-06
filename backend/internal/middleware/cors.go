package middleware

import (
	"net/http"
	"strings"
)

// CORS handles Cross-Origin Resource Sharing with strict origin matching and preflight support.
func CORS(allowedOrigins []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			var originAllowed bool
			for _, allowed := range allowedOrigins {
				if allowed == "*" || allowed == origin || strings.TrimRight(allowed, "/") == strings.TrimRight(origin, "/") {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")
					w.Header().Set("Access-Control-Expose-Headers", "X-Request-ID, Retry-After")
					w.Header().Set("Access-Control-Max-Age", "86400")
					originAllowed = true
					break
				}
			}

			// If it's a preflight OPTIONS request
			if r.Method == http.MethodOptions {
				if originAllowed {
					w.WriteHeader(http.StatusNoContent)
				} else {
					w.WriteHeader(http.StatusForbidden)
				}
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
