package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"strings"

	"portfolio-backend/internal/config"
	"portfolio-backend/internal/utils"
)

// CSRFProtection validates Origin/Referer headers and double-submit tokens on state-changing requests
func CSRFProtection(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Safe methods do not require CSRF verification
			if r.Method == http.MethodGet || r.Method == http.MethodHead || r.Method == http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			// Public unauthenticated contact form and login have their own rate-limiting & origin controls
			if r.URL.Path == "/api/contact" || r.URL.Path == "/api/admin/login" {
				next.ServeHTTP(w, r)
				return
			}

			// Only enforce CSRF for admin-authenticated routes
			if strings.HasPrefix(r.URL.Path, "/api/admin/") {
				// 1. Check Origin / Referer header against allowed origins and host
				origin := r.Header.Get("Origin")
				if origin == "" {
					referer := r.Header.Get("Referer")
					if referer != "" {
						if u, err := url.Parse(referer); err == nil {
							origin = u.Scheme + "://" + u.Host
						}
					}
				}

				if origin != "" {
					allowed := false
					for _, o := range cfg.AllowedOrigins {
						if o == "*" || o == origin || strings.TrimRight(o, "/") == strings.TrimRight(origin, "/") {
							allowed = true
							break
						}
					}
					if !allowed && cfg.Environment == "production" {
						utils.SendError(w, http.StatusForbidden, utils.CodeForbidden, "Cross-Origin request blocked by CSRF policy")
						return
					}
				}

				// 2. Double-submit token check (if cookie is present)
				csrfCookie, err := r.Cookie("portfolio_csrf")
				if err == nil && csrfCookie.Value != "" {
					headerToken := r.Header.Get("X-CSRF-Token")
					if headerToken == "" || headerToken != csrfCookie.Value {
						utils.SendError(w, http.StatusForbidden, utils.CodeCSRFInvalid, "Invalid or missing CSRF token")
						return
					}
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

// GenerateCSRFToken generates a signed CSRF token
func GenerateCSRFToken(sessionID, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(sessionID))
	return hex.EncodeToString(mac.Sum(nil))
}
