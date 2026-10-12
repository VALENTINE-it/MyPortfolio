package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"

	"portfolio-backend/internal/database/migrations"
)

// InitDB initializes SQLite connection, configures connection pool and pragmas,
// executes database migrations, and seeds initial portfolio data idempotently.
func InitDB(dataSourceName string) (*sql.DB, error) {
	if dataSourceName == "" {
		dataSourceName = os.Getenv("DATABASE_URL")
	}
	if dataSourceName == "" {
		dataSourceName = "./portfolio.db"
	}

	// Ensure parent directory exists if a path is provided
	dir := filepath.Dir(dataSourceName)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create database directory: %w", err)
		}
	}

	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure SQLite connection pool
	// In WAL mode, SQLite supports multiple concurrent readers and a serialized writer.
	// 25 max open connections prevents file descriptor exhaustion while maintaining high throughput.
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetConnMaxIdleTime(15 * time.Minute)

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Apply connection-level and file-level PRAGMAs for performance, concurrency, and durability
	pragmas := []string{
		"PRAGMA journal_mode = WAL;",
		"PRAGMA foreign_keys = ON;",
		"PRAGMA busy_timeout = 5000;",
		"PRAGMA synchronous = NORMAL;",
		"PRAGMA cache_size = -20000;", // ~20MB memory cache
		"PRAGMA temp_store = MEMORY;",
	}
	for _, pragma := range pragmas {
		if _, err := db.Exec(pragma); err != nil {
			slog.Warn("Error setting database pragma", "pragma", pragma, "error", err)
		}
	}

	// Run versioned migrations
	if err := migrations.Run(db); err != nil {
		return nil, fmt.Errorf("database migration failed: %w", err)
	}

	// Seed initial portfolio records if empty
	if err := seedInitialData(ctx, db); err != nil {
		slog.Warn("Error seeding initial portfolio data", "error", err)
	}

	return db, nil
}

// CheckIntegrity runs SQLite PRAGMA integrity_check to verify database health
func CheckIntegrity(db *sql.DB) error {
	var result string
	row := db.QueryRow("PRAGMA integrity_check;")
	if err := row.Scan(&result); err != nil {
		return fmt.Errorf("failed to run integrity check: %w", err)
	}
	if result != "ok" {
		return fmt.Errorf("database integrity check reported issues: %s", result)
	}
	return nil
}

// Backup creates an online atomic snapshot of the SQLite database using VACUUM INTO
func Backup(db *sql.DB, targetPath string) error {
	dir := filepath.Dir(targetPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create backup destination directory: %w", err)
	}

	// If file exists at targetPath, remove it before VACUUM INTO
	_ = os.Remove(targetPath)

	query := fmt.Sprintf("VACUUM INTO '%s';", targetPath)
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("failed to execute VACUUM INTO backup: %w", err)
	}

	slog.Info("Database backup completed successfully", "target", targetPath)
	return nil
}

// seedInitialData inserts starting projects and skills if tables are currently empty
func seedInitialData(ctx context.Context, db *sql.DB) error {
	// 1. Seed Projects if empty
	var projectCount int
	if err := db.QueryRow("SELECT COUNT(*) FROM projects").Scan(&projectCount); err != nil {
		return err
	}

	if projectCount == 0 {
		projects := []struct {
			Title        string
			Description  string
			Category     string
			Year         int
			Technologies string
			Image        string
			GithubURL    string
			LiveURL      string
		}{
			{
				Title:        "Safeguarding Reporting Platform",
				Description:  "A secure, confidential reporting platform designed for anonymous safeguarding disclosures, encrypted submission workflows, and case audit records.",
				Category:     "Full-Stack Web Application",
				Year:         2026,
				Technologies: `["React", "JavaScript", "Go", "SQLite", "REST APIs", "Security"]`,
				Image:        "/images/projects/safeguarding.svg",
				GithubURL:    "https://github.com/VALENTINE-it/safeguarding-platform",
				LiveURL:      "https://safeguarding.example.com",
			},
			{
				Title:        "Career Guidance Platform",
				Description:  "An educational guidance system that assesses academic performance and technical interests to recommend career pathways, skill milestones, and mentorship.",
				Category:     "Education Technology",
				Year:         2026,
				Technologies: `["React", "JavaScript", "Go", "REST APIs", "SQLite"]`,
				Image:        "/images/projects/career-guidance.svg",
				GithubURL:    "https://github.com/VALENTINE-it/career-guidance",
				LiveURL:      "",
			},
			{
				Title:        "Personal Developer Portfolio",
				Description:  "A digital editorial portfolio and professional showcase engineered with React and Go, featuring responsive fluid typography, clean REST architecture, and SQLite persistence.",
				Category:     "Full-Stack Portfolio",
				Year:         2026,
				Technologies: `["React", "Vite", "JavaScript", "Go", "REST API", "SQLite"]`,
				Image:        "/images/projects/portfolio.svg",
				GithubURL:    "https://github.com/VALENTINE-it/MyPortfolio",
				LiveURL:      "",
			},
		}

		stmt, err := db.Prepare(`
			INSERT INTO projects (title, description, category, year, technologies, image, github_url, live_url)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, p := range projects {
			if _, err := stmt.Exec(p.Title, p.Description, p.Category, p.Year, p.Technologies, p.Image, p.GithubURL, p.LiveURL); err != nil {
				return err
			}
		}
		slog.Info("Database: Successfully seeded initial projects")
	}

	// 2. Seed Skills if empty
	var skillCount int
	if err := db.QueryRow("SELECT COUNT(*) FROM skills").Scan(&skillCount); err != nil {
		return err
	}

	if skillCount == 0 {
		skills := []struct {
			Name     string
			Category string
		}{
			// Frontend
			{"HTML", "Frontend"},
			{"CSS", "Frontend"},
			{"JavaScript", "Frontend"},
			{"React", "Frontend"},
			{"Vite", "Frontend"},
			{"Responsive Design", "Frontend"},

			// Backend
			{"Go", "Backend"},
			{"REST APIs", "Backend"},
			{"SQLite", "Backend"},
			{"MongoDB", "Backend"},

			// Tools
			{"Git", "Tools"},
			{"GitHub", "Tools"},
			{"Linux", "Tools"},
			{"VS Code", "Tools"},

			// Blockchain
			{"Smart Contracts", "Blockchain Technology Developer"},
			{"Distributed Ledgers", "Blockchain Technology Developer"},
			{"Web3 Architecture", "Blockchain Technology Developer"},
			{"Cryptographic Protocols", "Blockchain Technology Developer"},

			// Other
			{"Networking", "Other"},
			{"Debugging", "Other"},
			{"Problem Solving", "Other"},
			{"Analytical Thinking", "Other"},
		}

		stmt, err := db.Prepare("INSERT INTO skills (name, category) VALUES (?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, s := range skills {
			if _, err := stmt.Exec(s.Name, s.Category); err != nil {
				return err
			}
		}
		slog.Info("Database: Successfully seeded initial skills")
	}

	return nil
}
