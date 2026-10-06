package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

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
	contactHandler := NewContactHandler(contactService)

	cleanup := func() {
		db.Close()
		os.RemoveAll(tempDir)
	}

	return projectHandler, skillHandler, contactHandler, cleanup
}

func TestHealthHandler(t *testing.T) {
	// 1. GET /api/health -> 200 {"status": "ok"}
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rr := httptest.NewRecorder()
	HandleHealth(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode health response: %v", err)
	}
	if resp["status"] != "ok" {
		t.Errorf("expected status 'ok', got '%s'", resp["status"])
	}

	// 2. Non-GET -> 405 Method Not Allowed
	postReq := httptest.NewRequest(http.MethodPost, "/api/health", nil)
	postRr := httptest.NewRecorder()
	HandleHealth(postRr, postReq)
	if postRr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405 on POST /api/health, got %d", postRr.Code)
	}
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
	req400 := httptest.NewRequest(http.MethodGet, "/api/projects/invalid-id", nil)
	rr400 := httptest.NewRecorder()
	projH.HandleProjects(rr400, req400)

	if rr400.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr400.Code)
	}

	// 4. SQL Injection attempt in ID
	reqSQLi := httptest.NewRequest(http.MethodGet, "/api/projects/1%20OR%201=1", nil)
	rrSQLi := httptest.NewRecorder()
	projH.HandleProjects(rrSQLi, reqSQLi)

	if rrSQLi.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 on SQL injection attempt, got %d", rrSQLi.Code)
	}
}

func TestSkillHandler_GetAll(t *testing.T) {
	_, skillH, _, cleanup := setupTestEnvironment(t)
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

	// 2. Grouped=false (flat list)
	reqFlat := httptest.NewRequest(http.MethodGet, "/api/skills?grouped=false", nil)
	rrFlat := httptest.NewRecorder()
	skillH.HandleSkills(rrFlat, reqFlat)

	if rrFlat.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rrFlat.Code)
	}

	var respFlat struct {
		Success bool           `json:"success"`
		Data    []models.Skill `json:"data"`
	}
	if err := json.NewDecoder(rrFlat.Body).Decode(&respFlat); err != nil {
		t.Fatalf("failed to decode flat skills response: %v", err)
	}
	if len(respFlat.Data) < 10 {
		t.Errorf("expected at least 10 skills in flat list, got %d", len(respFlat.Data))
	}

	// 3. Filter by category
	reqCat := httptest.NewRequest(http.MethodGet, "/api/skills?category=Frontend", nil)
	rrCat := httptest.NewRecorder()
	skillH.HandleSkills(rrCat, reqCat)

	if rrCat.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rrCat.Code)
	}
}

func TestContactHandler_ValidationAndSecurity(t *testing.T) {
	_, _, contactH, cleanup := setupTestEnvironment(t)
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
		"honeypot": "i-am-a-bot-filling-hidden-fields",
	}
	botBody, _ := json.Marshal(botPayload)
	botReq := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(botBody))
	botReq.Header.Set("Content-Type", "application/json")
	botRr := httptest.NewRecorder()
	contactH.HandleContact(botRr, botReq)

	if botRr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for honeypot triggered submission, got %d", botRr.Code)
	}

	// 3. DisallowUnknownFields test -> unknown JSON field rejected
	unknownFieldPayload := `{"name":"Valentine","email":"v@test.com","subject":"Hi","message":"Test","extra_field":"unexpected"}`
	unknownReq := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(unknownFieldPayload))
	unknownReq.Header.Set("Content-Type", "application/json")
	unknownRr := httptest.NewRecorder()
	contactH.HandleContact(unknownRr, unknownReq)

	if unknownRr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for unknown JSON field, got %d", unknownRr.Code)
	}

	// 4. Multiple JSON objects rejected
	multiJSONPayload := `{"name":"A","email":"a@test.com","subject":"S","message":"M"}{"name":"B","email":"b@test.com","subject":"S","message":"M"}`
	multiReq := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(multiJSONPayload))
	multiReq.Header.Set("Content-Type", "application/json")
	multiRr := httptest.NewRecorder()
	contactH.HandleContact(multiRr, multiReq)

	if multiRr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for multiple JSON objects in body, got %d", multiRr.Code)
	}

	// 5. Message too long (>5000 characters)
	longMsg := strings.Repeat("A", 5001)
	longPayload := map[string]string{
		"name":    "Valentine",
		"email":   "test@example.com",
		"subject": "Long message",
		"message": longMsg,
	}
	longBody, _ := json.Marshal(longPayload)
	longReq := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(longBody))
	longReq.Header.Set("Content-Type", "application/json")
	longRr := httptest.NewRecorder()
	contactH.HandleContact(longRr, longReq)

	if longRr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for oversized message, got %d", longRr.Code)
	}
}
