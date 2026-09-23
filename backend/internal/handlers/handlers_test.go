package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"portfolio-backend/internal/database"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
)

func setupTestEnvironment(t *testing.T) (*ProjectHandler, *SkillHandler, *ContactHandler, func()) {
	tempDir, err := os.MkdirTemp("", "portfolio_test_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	dbPath := filepath.Join(tempDir, "test.db")
	db, err := database.InitDB(dbPath)
	if err != nil {
		t.Fatalf("failed to init db: %v", err)
	}

	projectRepo := repositories.NewProjectRepository(db)
	skillRepo := repositories.NewSkillRepository(db)
	contactRepo := repositories.NewContactRepository(db)

	projectService := services.NewProjectService(projectRepo)
	skillService := services.NewSkillService(skillRepo)
	contactService := services.NewContactService(contactRepo)

	projectHandler := NewProjectHandler(projectService)
	skillHandler := NewSkillHandler(skillService)
	// For testing, use a small window for the rate limiter
	contactHandler := &ContactHandler{
		service:     contactService,
		rateLimiter: NewRateLimiter(3, 1*time.Second),
	}

	cleanup := func() {
		db.Close()
		os.RemoveAll(tempDir)
	}

	return projectHandler, skillHandler, contactHandler, cleanup
}

func TestProjectHandler_GetAll(t *testing.T) {
	projH, _, _, cleanup := setupTestEnvironment(t)
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
	if len(resp.Data[0].Technologies) == 0 {
		t.Errorf("expected technologies array to be populated")
	}
}

func TestProjectHandler_GetByID(t *testing.T) {
	projH, _, _, cleanup := setupTestEnvironment(t)
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
	req400 := httptest.NewRequest(http.MethodGet, "/api/projects/invalid", nil)
	rr400 := httptest.NewRecorder()
	projH.HandleProjects(rr400, req400)

	if rr400.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr400.Code)
	}
}

func TestSkillHandler_GetAll(t *testing.T) {
	_, skillH, _, cleanup := setupTestEnvironment(t)
	defer cleanup()

	req := httptest.NewRequest(http.MethodGet, "/api/skills", nil)
	rr := httptest.NewRecorder()

	skillH.HandleSkills(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp struct {
		Success bool                      `json:"success"`
		Data    []models.SkillsByCategory `json:"data"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !resp.Success {
		t.Errorf("expected success true")
	}
	if len(resp.Data) == 0 {
		t.Errorf("expected grouped categories")
	}
}

func TestContactHandler_Submit(t *testing.T) {
	_, _, contactH, cleanup := setupTestEnvironment(t)
	defer cleanup()

	// 1. Successful submission
	payload := map[string]string{
		"name":    "Valentine Test",
		"email":   "test@example.com",
		"subject": "Collaboration Opportunity",
		"message": "Hello Valentine, this is a test message regarding a project.",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	contactH.HandleContact(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d: %s", rr.Code, rr.Body.String())
	}

	// 2. Validation failure - invalid email
	badPayload := map[string]string{
		"name":    "Valentine Test",
		"email":   "not-an-email",
		"subject": "Hello",
		"message": "Valid test message here",
	}
	badBody, _ := json.Marshal(badPayload)

	badReq := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(badBody))
	badReq.Header.Set("Content-Type", "application/json")
	badRr := httptest.NewRecorder()

	contactH.HandleContact(badRr, badReq)

	if badRr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", badRr.Code)
	}

	// 3. Method not allowed
	getReq := httptest.NewRequest(http.MethodGet, "/api/contact", nil)
	getRr := httptest.NewRecorder()
	contactH.HandleContact(getRr, getReq)

	if getRr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", getRr.Code)
	}
}
