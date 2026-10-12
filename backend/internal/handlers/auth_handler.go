package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"portfolio-backend/internal/config"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"
	"portfolio-backend/internal/utils"
)

type AuthHandler struct {
	authService *services.AuthService
	cfg         *config.Config
}

func NewAuthHandler(authService *services.AuthService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		cfg:         cfg,
	}
}

// HandleLogin handles POST /api/admin/login
func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.AdminLoginRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Invalid JSON payload")
		return
	}

	if req.Username == "" || req.Password == "" {
		utils.SendError(w, http.StatusBadRequest, utils.CodeBadRequest, "Username and password are required")
		return
	}

	userAgent := r.UserAgent()
	ipAddress := r.RemoteAddr

	session, admin, err := h.authService.Login(r.Context(), req.Username, req.Password, userAgent, ipAddress)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, utils.CodeUnauthorized, "Invalid username or password")
		return
	}

	// Set HttpOnly, Secure session cookie
	sameSite := http.SameSiteLaxMode
	if stringsEqualFold(h.cfg.CookieSameSite, "strict") {
		sameSite = http.SameSiteStrictMode
	} else if stringsEqualFold(h.cfg.CookieSameSite, "none") {
		sameSite = http.SameSiteNoneMode
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "portfolio_session",
		Value:    session.ID,
		Path:     "/",
		Domain:   h.cfg.CookieDomain,
		Expires:  session.ExpiresAt,
		MaxAge:   int(h.cfg.SessionTTL.Seconds()),
		HttpOnly: true,
		Secure:   h.cfg.CookieSecure,
		SameSite: sameSite,
	})

	// Generate and set CSRF cookie
	csrfToken := middleware.GenerateCSRFToken(session.ID, h.cfg.CSRFSecret)
	http.SetCookie(w, &http.Cookie{
		Name:     "portfolio_csrf",
		Value:    csrfToken,
		Path:     "/",
		Domain:   h.cfg.CookieDomain,
		Expires:  session.ExpiresAt,
		MaxAge:   int(h.cfg.SessionTTL.Seconds()),
		HttpOnly: false, // Accessible by JavaScript client to send in X-CSRF-Token header
		Secure:   h.cfg.CookieSecure,
		SameSite: sameSite,
	})

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Login successful",
		"data": map[string]interface{}{
			"token": session.ID,
			"admin": models.AdminProfile{
				ID:       admin.ID,
				Username: admin.Username,
			},
			"csrfToken": csrfToken,
			"expiresAt": session.ExpiresAt.Format(time.RFC3339),
		},
	})
}

// HandleLogout handles POST /api/admin/logout
func (h *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}

	session := middleware.GetSessionFromContext(r.Context())
	if session != nil {
		_ = h.authService.Logout(r.Context(), session.ID)
	}

	// Invalidate session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "portfolio_session",
		Value:    "",
		Path:     "/",
		Domain:   h.cfg.CookieDomain,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   h.cfg.CookieSecure,
	})

	// Invalidate CSRF cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "portfolio_csrf",
		Value:    "",
		Path:     "/",
		Domain:   h.cfg.CookieDomain,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: false,
		Secure:   h.cfg.CookieSecure,
	})

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Logged out successfully",
	})
}

// HandleMe handles GET /api/admin/me
func (h *AuthHandler) HandleMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}

	admin := middleware.GetAdminFromContext(r.Context())
	if admin == nil {
		utils.SendError(w, http.StatusUnauthorized, utils.CodeUnauthorized, "Not authenticated")
		return
	}

	utils.SendSuccess(w, http.StatusOK, models.AdminProfile{
		ID:       admin.ID,
		Username: admin.Username,
	})
}

// HandleCSRF handles GET /api/admin/csrf
func (h *AuthHandler) HandleCSRF(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, utils.CodeMethodNotAllowed, "Method not allowed")
		return
	}

	session := middleware.GetSessionFromContext(r.Context())
	sessionID := "anonymous"
	if session != nil {
		sessionID = session.ID
	}

	token := middleware.GenerateCSRFToken(sessionID, h.cfg.CSRFSecret)

	http.SetCookie(w, &http.Cookie{
		Name:     "portfolio_csrf",
		Value:    token,
		Path:     "/",
		Domain:   h.cfg.CookieDomain,
		HttpOnly: false,
		Secure:   h.cfg.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	})

	utils.SendSuccess(w, http.StatusOK, map[string]string{
		"csrfToken": token,
	})
}

func stringsEqualFold(s1, s2 string) bool {
	return len(s1) == len(s2) && (s1 == s2 || (len(s1) > 0 && s1[0] == s2[0]))
}
