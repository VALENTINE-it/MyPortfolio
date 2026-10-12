package middleware

import (
	"context"
	"net/http"
	"strings"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"
	"portfolio-backend/internal/utils"
)

type authContextKey string

const (
	AdminContextKey   authContextKey = "auth_admin"
	SessionContextKey authContextKey = "auth_session"
)

// Authenticate checks for session cookie or Authorization header and attaches the admin to context
func Authenticate(authService *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var sessionID string

			// 1. Try reading cookie
			cookie, err := r.Cookie("portfolio_session")
			if err == nil && cookie.Value != "" {
				sessionID = cookie.Value
			}

			// 2. Fall back to Authorization: Bearer <token>
			if sessionID == "" {
				authHeader := r.Header.Get("Authorization")
				if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
					sessionID = strings.TrimSpace(authHeader[7:])
				}
			}

			if sessionID != "" {
				admin, session, err := authService.ValidateSession(r.Context(), sessionID)
				if err == nil && admin != nil {
					ctx := context.WithValue(r.Context(), AdminContextKey, admin)
					ctx = context.WithValue(ctx, SessionContextKey, session)
					r = r.WithContext(ctx)
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAdmin rejects unauthenticated requests with HTTP 401 Unauthorized
func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		admin := GetAdminFromContext(r.Context())
		if admin == nil {
			utils.SendError(w, http.StatusUnauthorized, utils.CodeUnauthorized, "Authentication required to access this resource")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// GetAdminFromContext extracts the authenticated admin from context
func GetAdminFromContext(ctx context.Context) *models.Admin {
	if ctx == nil {
		return nil
	}
	if a, ok := ctx.Value(AdminContextKey).(*models.Admin); ok {
		return a
	}
	return nil
}

// GetSessionFromContext extracts the authenticated session from context
func GetSessionFromContext(ctx context.Context) *models.Session {
	if ctx == nil {
		return nil
	}
	if s, ok := ctx.Value(SessionContextKey).(*models.Session); ok {
		return s
	}
	return nil
}
