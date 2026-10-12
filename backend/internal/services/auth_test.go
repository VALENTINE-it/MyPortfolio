package services

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"portfolio-backend/internal/database"
	"portfolio-backend/internal/repositories"
)

func TestAuthService_FullLifecycle(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "auth_test_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	db, err := database.InitDB(filepath.Join(tempDir, "test.db"))
	if err != nil {
		t.Fatalf("failed to init db: %v", err)
	}
	defer db.Close()

	adminRepo := repositories.NewAdminRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)
	authService := NewAuthService(adminRepo, sessionRepo, 24*time.Hour)

	ctx := context.Background()

	// 1. Password hashing & verification
	hash, err := authService.HashPassword("validPassword123!")
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}
	if !authService.VerifyPassword("validPassword123!", hash) {
		t.Fatalf("expected password verification to succeed")
	}
	if authService.VerifyPassword("wrongPassword", hash) {
		t.Fatalf("expected incorrect password to fail")
	}

	// 2. Create Admin
	admin, err := authService.CreateAdmin(ctx, "admin_user", "validPassword123!")
	if err != nil {
		t.Fatalf("failed to create admin: %v", err)
	}
	if admin.Username != "admin_user" {
		t.Fatalf("expected username admin_user, got %s", admin.Username)
	}

	// 3. Duplicate Admin Creation Should Fail
	_, err = authService.CreateAdmin(ctx, "admin_user", "anotherPassword123!")
	if err == nil {
		t.Fatalf("expected error on duplicate username, got nil")
	}

	// 4. Successful Login
	session, loggedInAdmin, err := authService.Login(ctx, "admin_user", "validPassword123!", "Go-Test-Agent", "127.0.0.1")
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}
	if session.ID == "" || loggedInAdmin.ID != admin.ID {
		t.Fatalf("invalid session or admin return")
	}

	// 5. Session Validation
	validatedAdmin, validatedSession, err := authService.ValidateSession(ctx, session.ID)
	if err != nil {
		t.Fatalf("session validation failed: %v", err)
	}
	if validatedAdmin.ID != admin.ID || validatedSession.ID != session.ID {
		t.Fatalf("validated session mismatch")
	}

	// 6. Logout
	if err := authService.Logout(ctx, session.ID); err != nil {
		t.Fatalf("logout failed: %v", err)
	}

	// 7. Validation After Logout Should Fail
	_, _, err = authService.ValidateSession(ctx, session.ID)
	if err == nil {
		t.Fatalf("expected error validating destroyed session, got nil")
	}
}
