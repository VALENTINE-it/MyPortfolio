package middleware

import (
	"net/http"
)

// SecurityHeaders applies essential security headers to every response.
func SecurityHeaders(isProduction bool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevent MIME-sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Prevent clickjacking / iframe embedding
		w.Header().Set("X-Frame-Options", "DENY")

		// Restrict referrer leakage
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		// Restrict browser hardware permissions
		w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=(), payment=()")

		// Content Security Policy for API endpoints
		w.Header().Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none';")

		// Strict-Transport-Security (only in production over HTTPS)
		if isProduction {
			w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		}

		next.ServeHTTP(w, r)
	})
}
