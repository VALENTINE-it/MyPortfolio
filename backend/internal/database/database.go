package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// InitDB initializes SQLite connection, applies pragmas, creates tables, and seeds initial data.
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

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Pragmas for performance, durability, and foreign key enforcement
	pragmas := []string{
		"PRAGMA journal_mode = WAL;",
		"PRAGMA foreign_keys = ON;",
		"PRAGMA busy_timeout = 5000;",
		"PRAGMA synchronous = NORMAL;",
	}
	for _, pragma := range pragmas {
		if _, err := db.Exec(pragma); err != nil {
			log.Printf("Warning: error setting pragma %q: %v", pragma, err)
		}
	}

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	if err := seedInitialData(db); err != nil {
		log.Printf("Warning: error seeding initial data: %v", err)
	}

	return db, nil
}

// createTables executes DDL for projects, skills, contacts, and performance indexes.
func createTables(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		category TEXT NOT NULL,
		year INTEGER NOT NULL,
		technologies TEXT NOT NULL,
		image TEXT NOT NULL,
		github_url TEXT,
		live_url TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_projects_year ON projects(year DESC);

	CREATE TABLE IF NOT EXISTS skills (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		category TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_skills_category ON skills(category);

	CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		subject TEXT NOT NULL,
		message TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_contacts_created_at ON contacts(created_at DESC);
	`
	_, err := db.Exec(schema)
	return err
}

// seedInitialData inserts starting projects and skills if tables are currently empty.
func seedInitialData(db *sql.DB) error {
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
		log.Println("Database: Successfully seeded initial projects.")
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
		log.Println("Database: Successfully seeded initial skills.")
	}

	return nil
}
