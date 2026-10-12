package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/utils"
)

type AuthService struct {
	adminRepo   *repositories.AdminRepository
	sessionRepo *repositories.SessionRepository
	sessionTTL  time.Duration
}

func NewAuthService(adminRepo *repositories.AdminRepository, sessionRepo *repositories.SessionRepository, sessionTTL time.Duration) *AuthService {
	if sessionTTL <= 0 {
		sessionTTL = 24 * time.Hour
	}
	return &AuthService{
		adminRepo:   adminRepo,
		sessionRepo: sessionRepo,
		sessionTTL:  sessionTTL,
	}
}

// HashPassword hashes a raw password string using bcrypt with cost 12
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

// VerifyPassword compares a plaintext password with a hashed password
func (s *AuthService) VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateSessionID generates 32 cryptographically secure random bytes (64 hex characters)
func (s *AuthService) GenerateSessionID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate secure session ID: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

// Login authenticates an admin by username and password, rotates and returns a new session
func (s *AuthService) Login(ctx context.Context, username, password, userAgent, ipAddress string) (*models.Session, *models.Admin, error) {
	admin, err := s.adminRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, repositories.ErrAdminNotFound) {
			slog.Warn("Login failed: username not found", "username", username, "ip", ipAddress)
			return nil, nil, utils.ErrInvalidCredentials
		}
		return nil, nil, err
	}

	if !s.VerifyPassword(password, admin.PasswordHash) {
		slog.Warn("Login failed: password mismatch", "username", username, "ip", ipAddress)
		return nil, nil, utils.ErrInvalidCredentials
	}

	// Generate secure session
	sessionID, err := s.GenerateSessionID()
	if err != nil {
		return nil, nil, err
	}

	now := time.Now().UTC()
	session := &models.Session{
		ID:        sessionID,
		AdminID:   admin.ID,
		ExpiresAt: now.Add(s.sessionTTL),
		CreatedAt: now,
		UserAgent: userAgent,
		IPAddress: ipAddress,
	}

	if err := s.sessionRepo.CreateSession(ctx, session); err != nil {
		return nil, nil, fmt.Errorf("failed to create session: %w", err)
	}

	slog.Info("Admin login successful", "admin_id", admin.ID, "username", admin.Username, "ip", ipAddress)
	return session, admin, nil
}

// ValidateSession validates that a session token is active and not expired
func (s *AuthService) ValidateSession(ctx context.Context, sessionID string) (*models.Admin, *models.Session, error) {
	if sessionID == "" {
		return nil, nil, utils.ErrUnauthorized
	}

	session, err := s.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		if errors.Is(err, repositories.ErrSessionNotFound) {
			return nil, nil, utils.ErrUnauthorized
		}
		return nil, nil, err
	}

	if time.Now().UTC().After(session.ExpiresAt) {
		_ = s.sessionRepo.DeleteSession(ctx, sessionID)
		return nil, nil, utils.ErrSessionExpired
	}

	admin, err := s.adminRepo.GetByID(ctx, session.AdminID)
	if err != nil {
		return nil, nil, utils.ErrUnauthorized
	}

	return admin, session, nil
}

// Logout destroys an active session
func (s *AuthService) Logout(ctx context.Context, sessionID string) error {
	if sessionID == "" {
		return nil
	}
	return s.sessionRepo.DeleteSession(ctx, sessionID)
}

// CreateAdmin creates an administrator account (used by secure CLI provisioning)
func (s *AuthService) CreateAdmin(ctx context.Context, username, password string) (*models.Admin, error) {
	if !utils.ValidateUsername(username) {
		return nil, fmt.Errorf("invalid username format: must be 3-32 alphanumeric characters, dashes or underscores")
	}
	if len(password) < 10 {
		return nil, fmt.Errorf("password must be at least 10 characters")
	}

	existing, _ := s.adminRepo.GetByUsername(ctx, username)
	if existing != nil {
		return nil, fmt.Errorf("an administrator with username %q already exists", username)
	}

	hash, err := s.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return s.adminRepo.Create(ctx, username, hash)
}

// CleanupExpiredSessions prunes expired sessions
func (s *AuthService) CleanupExpiredSessions(ctx context.Context) (int64, error) {
	return s.sessionRepo.DeleteExpiredSessions(ctx)
}
