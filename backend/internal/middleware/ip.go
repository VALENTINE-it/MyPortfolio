package middleware

import (
	"net"
	"net/http"
	"strings"
)

// GetClientIP extracts the client IP address.
// By default, it uses RemoteAddr to prevent header spoofing attacks.
// It ONLY checks X-Forwarded-For or X-Real-IP if trustProxy is explicitly enabled.
func GetClientIP(r *http.Request, trustProxy bool) string {
	if trustProxy {
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			parts := strings.Split(xff, ",")
			if len(parts) > 0 {
				candidate := strings.TrimSpace(parts[0])
				if ip := net.ParseIP(candidate); ip != nil {
					return candidate
				}
			}
		}
		if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
			candidate := strings.TrimSpace(xrip)
			if ip := net.ParseIP(candidate); ip != nil {
				return candidate
			}
		}
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// Could be an IP without port
		if ip := net.ParseIP(r.RemoteAddr); ip != nil {
			return r.RemoteAddr
		}
		return "127.0.0.1"
	}
	return host
}
