package database

import (
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
}
