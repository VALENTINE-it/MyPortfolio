package database

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestInitDB(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "portfolio_db_test_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	dbPath := filepath.Join(tempDir, "test.db")

	db, err := InitDB(dbPath)
	if err != nil {
		t.Fatalf("failed to initialize db: %v", err)
	}
	defer db.Close()

	// Verify tables exist
	tables := []string{"projects", "skills", "contacts"}
	for _, table := range tables {
		var name string
		err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name=?", table).Scan(&name)
		if err != nil {
			t.Errorf("expected table %s to exist, error: %v", table, err)
		}
	}

	// Verify indexes exist
	indexes := []string{"idx_projects_year", "idx_skills_category", "idx_contacts_created_at"}
	for _, idx := range indexes {
		var name string
		err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='index' AND name=?", idx).Scan(&name)
		if err != nil {
			t.Errorf("expected index %s to exist, error: %v", idx, err)
		}
	}

	// Verify foreign_keys pragma is ON
	var fk int
	if err := db.QueryRow("PRAGMA foreign_keys").Scan(&fk); err != nil {
		t.Fatalf("failed to query foreign_keys pragma: %v", err)
	}
	if fk != 1 {
		t.Errorf("expected foreign_keys to be enabled (1), got %d", fk)
	}

	// Verify busy_timeout pragma
	var busyTimeout int
	if err := db.QueryRow("PRAGMA busy_timeout").Scan(&busyTimeout); err != nil {
		t.Fatalf("failed to query busy_timeout pragma: %v", err)
	}
	if busyTimeout < 5000 {
		t.Errorf("expected busy_timeout to be at least 5000, got %d", busyTimeout)
	}

	// Verify seeding works
	var projectCount int
	if err := db.QueryRow("SELECT COUNT(*) FROM projects").Scan(&projectCount); err != nil {
		t.Fatalf("failed to count projects: %v", err)
	}
	if projectCount < 3 {
		t.Errorf("expected at least 3 seeded projects, got %d", projectCount)
	}

	var skillCount int
	if err := db.QueryRow("SELECT COUNT(*) FROM skills").Scan(&skillCount); err != nil {
		t.Fatalf("failed to count skills: %v", err)
	}
	if skillCount < 10 {
		t.Errorf("expected at least 10 seeded skills, got %d", skillCount)
	}

	// Verify idempotency of seedInitialData
	if err := seedInitialData(context.Background(), db); err != nil {
		t.Fatalf("second seedInitialData call failed: %v", err)
	}
	var newProjectCount int
	_ = db.QueryRow("SELECT COUNT(*) FROM projects").Scan(&newProjectCount)
	if newProjectCount != projectCount {
		t.Errorf("expected project count to remain %d after second seed, got %d", projectCount, newProjectCount)
	}
}

func TestContextCancellation(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "portfolio_ctx_test_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	dbPath := filepath.Join(tempDir, "test.db")
	db, err := InitDB(dbPath)
	if err != nil {
		t.Fatalf("failed to init db: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err = db.QueryContext(ctx, "SELECT COUNT(*) FROM projects")
	if err == nil {
		t.Errorf("expected error on cancelled context query, got nil")
	}
}
