package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"portfolio-backend/internal/config"
	"portfolio-backend/internal/database"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/monitoring"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
)

func setupTestEnvironment(t *testing.T) (*ProjectHandler, *SkillHandler, *ContactHandler, *AuthHandler, *services.AuthService, func()) {
	tempDir, err := os.MkdirTemp("", "portfolio_test_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	dbPath := filepath.Join(tempDir, "test.db")
	db, err := database.InitDB(dbPath)
	if err != nil {
		t.Fatalf("failed to init db: %v", err)
	}

	cfg, _ := config.LoadConfig()

	projectRepo := repositories.NewProjectRepository(db)
	skillRepo := repositories.NewSkillRepository(db)
	contactRepo := repositories.NewContactRepository(db)
	adminRepo := repositories.NewAdminRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)

	projectService := services.NewProjectService(projectRepo)
	skillService := services.NewSkillService(skillRepo)
	contactService := services.NewContactService(contactRepo)
	authService := services.NewAuthService(adminRepo, sessionRepo, 24*time.Hour)

	projectHandler := NewProjectHandler(projectService, nil)
	skillHandler := NewSkillHandler(skillService, nil)
	contactHandler := NewContactHandler(contactService)
	authHandler := NewAuthHandler(authService, cfg)

	cleanup := func() {
		db.Close()
		os.RemoveAll(tempDir)
	}

	return projectHandler, skillHandler, contactHandler, authHandler, authService, cleanup
}

func TestHealthHandler(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "portfolio_health_test_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	db, err := database.InitDB(filepath.Join(tempDir, "test.db"))
	if err != nil {
		t.Fatalf("failed to init db: %v", err)
	}
	defer db.Close()

	healthH := monitoring.NewHealthHandler(db)

	// 1. GET /api/health -> 200 {"success": true, "message": "API is running"}
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rr := httptest.NewRecorder()
	healthH.LegacyHealthHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode health response: %v", err)
	}
	if resp["success"] != true {
		t.Errorf("expected success true, got %v", resp["success"])
	}

	// 2. GET /health/live
	liveReq := httptest.NewRequest(http.MethodGet, "/health/live", nil)
	liveRr := httptest.NewRecorder()
	healthH.LivenessHandler(liveRr, liveReq)
	if liveRr.Code != http.StatusOK {
		t.Errorf("expected 200 on /health/live, got %d", liveRr.Code)
	}

	// 3. GET /health/ready
	readyReq := httptest.NewRequest(http.MethodGet, "/health/ready", nil)
	readyRr := httptest.NewRecorder()
	healthH.ReadinessHandler(readyRr, readyReq)
	if readyRr.Code != http.StatusOK {
		t.Errorf("expected 200 on /health/ready, got %d", readyRr.Code)
	}
}

func TestProjectHandler_GetAll(t *testing.T) {
	projH, _, _, _, _, cleanup := setupTestEnvironment(t)
	defer cleanup()

	req := httptest.NewRequest(http.MethodGet, "/api/projects", nil)
	rr := httptest.NewRecorder()

	projH.HandleProjects(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp struct {
		Success bool             `json:"success"`
		Data    []models.Project `json:"data"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !resp.Success {
		t.Errorf("expected success true")
	}
	if len(resp.Data) < 3 {
		t.Errorf("expected at least 3 projects, got %d", len(resp.Data))
	}
}

func TestProjectHandler_GetByID(t *testing.T) {
	projH, _, _, _, _, cleanup := setupTestEnvironment(t)
	defer cleanup()

	// 1. Valid ID = 1
	req := httptest.NewRequest(http.MethodGet, "/api/projects/1", nil)
	rr := httptest.NewRecorder()
	projH.HandleProjects(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	// 2. Non-existent ID = 9999
	req404 := httptest.NewRequest(http.MethodGet, "/api/projects/9999", nil)
	rr404 := httptest.NewRecorder()
	projH.HandleProjects(rr404, req404)

	if rr404.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr404.Code)
	}

	// 3. Invalid ID format
	req400 := httptest.NewRequest(http.MethodGet, "/api/projects/invalid-id", nil)
	rr400 := httptest.NewRecorder()
	projH.HandleProjects(rr400, req400)

	if rr400.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr400.Code)
	}
}

func TestSkillHandler_GetAll(t *testing.T) {
	_, skillH, _, _, _, cleanup := setupTestEnvironment(t)
	defer cleanup()

	// 1. Grouped by default
	req := httptest.NewRequest(http.MethodGet, "/api/skills", nil)
	rr := httptest.NewRecorder()
	skillH.HandleSkills(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var respGrouped struct {
		Success bool                      `json:"success"`
		Data    []models.SkillsByCategory `json:"data"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&respGrouped); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if !respGrouped.Success || len(respGrouped.Data) == 0 {
		t.Errorf("expected grouped categories in response")
	}
}

func TestContactHandler_ValidationAndSecurity(t *testing.T) {
	_, _, contactH, _, _, cleanup := setupTestEnvironment(t)
	defer cleanup()

	// 1. Successful submission
	validPayload := map[string]string{
		"name":     "Valentine Test",
		"email":    "test@example.com",
		"subject":  "Collaboration Opportunity",
		"message":  "Hello Valentine, this is a test message regarding a project.",
		"honeypot": "",
	}
	body, _ := json.Marshal(validPayload)

	req := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	contactH.HandleContact(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d: %s", rr.Code, rr.Body.String())
	}

	// 2. Honeypot filled -> bot detection 400
	botPayload := map[string]string{
		"name":     "Bot Spammer",
		"email":    "bot@spam.com",
		"subject":  "Spam subject",
		"message":  "Spam message body",
		"honeypot": "i-am-a-bot",
	}
	botBody, _ := json.Marshal(botPayload)
	botReq := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(botBody))
	botReq.Header.Set("Content-Type", "application/json")
	botRr := httptest.NewRecorder()
	contactH.HandleContact(botRr, botReq)

	if botRr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for honeypot, got %d", botRr.Code)
	}
}

func TestAdminAuthFlow(t *testing.T) {
	projH, _, _, authH, authService, cleanup := setupTestEnvironment(t)
	defer cleanup()

	ctx := context.Background()

	// 1. Provision Admin
	admin, err := authService.CreateAdmin(ctx, "admin_test", "supersecret123")
	if err != nil {
		t.Fatalf("failed to create admin: %v", err)
	}

	// 2. Attempt login with bad password -> 401
	badLoginPayload := map[string]string{
		"username": "admin_test",
		"password": "wrongpassword",
	}
	badBody, _ := json.Marshal(badLoginPayload)
	badReq := httptest.NewRequest(http.MethodPost, "/api/admin/login", bytes.NewReader(badBody))
	badRr := httptest.NewRecorder()
	authH.HandleLogin(badRr, badReq)

	if badRr.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401 on bad password, got %d", badRr.Code)
	}

	// 3. Successful login -> 200 with cookie
	loginPayload := map[string]string{
		"username": "admin_test",
		"password": "supersecret123",
	}
	loginBody, _ := json.Marshal(loginPayload)
	loginReq := httptest.NewRequest(http.MethodPost, "/api/admin/login", bytes.NewReader(loginBody))
	loginRr := httptest.NewRecorder()
	authH.HandleLogin(loginRr, loginReq)

	if loginRr.Code != http.StatusOK {
		t.Fatalf("expected status 200 on valid login, got %d", loginRr.Code)
	}

	// Extract session cookie
	cookies := loginRr.Result().Cookies()
	var sessionCookie *http.Cookie
	for _, c := range cookies {
		if c.Name == "portfolio_session" {
			sessionCookie = c
			break
		}
	}
	if sessionCookie == nil {
		t.Fatalf("expected portfolio_session cookie in login response")
	}

	// 4. Test RequireAdmin middleware blocks unauthorized requests
	protectedHandler := middleware.RequireAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	unauthReq := httptest.NewRequest(http.MethodGet, "/api/admin/projects", nil)
	unauthRr := httptest.NewRecorder()
	protectedHandler.ServeHTTP(unauthRr, unauthReq)

	if unauthRr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for unauthenticated access, got %d", unauthRr.Code)
	}

	// 5. Test RequireAdmin allows authenticated request with session context
	authReq := httptest.NewRequest(http.MethodGet, "/api/admin/projects", nil)
	authReq = authReq.WithContext(context.WithValue(authReq.Context(), middleware.AdminContextKey, admin))
	authRr := httptest.NewRecorder()
	protectedHandler.ServeHTTP(authRr, authReq)

	if authRr.Code != http.StatusOK {
		t.Fatalf("expected 200 for authenticated access, got %d", authRr.Code)
	}

	// 6. Test Admin Project Creation
	newProj := models.ProjectInput{
		Title:        "Automated Test Project",
		Description:  "Description for automated test project",
		Category:     "Full-Stack",
		Year:         2026,
		Technologies: []string{"Go", "SQLite", "React"},
		Image:        "/images/test.svg",
		GithubURL:    "https://github.com/test",
		LiveURL:      "https://test.example.com",
	}
	newProjBody, _ := json.Marshal(newProj)
	createReq := httptest.NewRequest(http.MethodPost, "/api/admin/projects", bytes.NewReader(newProjBody))
	createRr := httptest.NewRecorder()
	projH.HandleAdminProjects(createRr, createReq)

	if createRr.Code != http.StatusCreated {
		t.Fatalf("expected 201 on project creation, got %d: %s", createRr.Code, createRr.Body.String())
	}
}
