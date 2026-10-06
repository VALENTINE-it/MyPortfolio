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
)

/*
==============================================================================
DATABASE BACKUP & RESTORE DOCUMENTATION
==============================================================================

1. Storage Location:
   By default, the SQLite database is stored at `./data/portfolio.db` (or via
   DATABASE_PATH). This directory is kept outside publicly accessible web roots.

2. Online Backup Process (Zero-Downtime Safe via WAL mode):
   Run SQLite's safe vacuum/backup command:
     sqlite3 ./data/portfolio.db "VACUUM INTO './data/backups/portfolio_backup_$(date +%Y%m%d_%H%M%S).db'"
   Or run a standard live snapshot copy:
     sqlite3 ./data/portfolio.db ".backup './data/backups/backup.db'"

3. Restoration Process:
   To restore from a backup file:
     1. Stop the backend service (SIGTERM/graceful shutdown).
     2. Copy backup file over existing:
        cp ./data/backups/portfolio_backup_YYYYMMDD.db ./data/portfolio.db
     3. Restart backend service.

4. Public Access Security:
   The .db, .db-wal, and .db-shm files must NEVER be served directly by any web server.
==============================================================================
*/

// InitDB initializes SQLite connection pool, applies security/performance pragmas,
// applies schema migrations, and seeds initial data if tables are empty.
func InitDB(dataSourceName string) (*sql.DB, error) {
	if dataSourceName == "" {
		dataSourceName = os.Getenv("DATABASE_PATH")
	}
	if dataSourceName == "" {
		dataSourceName = os.Getenv("DATABASE_URL")
	}
	if dataSourceName == "" {
		dataSourceName = "./data/portfolio.db"
	}

	// Ensure parent directory exists securely
	dir := filepath.Dir(dataSourceName)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0750); err != nil {
			return nil, fmt.Errorf("failed to create database directory: %w", err)
		}
	}

	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Connection Pool Limits: SQLite is a single-file database.
	// In WAL mode, SQLite allows concurrent readers and one writer.
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetConnMaxIdleTime(15 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Verify connection
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Pragmas for durability, concurrency, foreign key enforcement, and timeout resilience
	pragmas := []string{
		"PRAGMA journal_mode = WAL;",
		"PRAGMA foreign_keys = ON;",
		"PRAGMA busy_timeout = 5000;",
		"PRAGMA synchronous = NORMAL;",
		"PRAGMA temp_store = MEMORY;",
		"PRAGMA cache_size = -64000;", // 64MB cache
	}
	for _, pragma := range pragmas {
		if _, err := db.ExecContext(ctx, pragma); err != nil {
			slog.Warn("Database: pragma setting notice", "pragma", pragma, "error", err)
		}
	}

	// Apply migrations
	if err := createTables(ctx, db); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to apply migrations: %w", err)
	}

	// Seed initial data if necessary
	if err := seedInitialData(ctx, db); err != nil {
		slog.Warn("Database: notice during seed", "error", err)
	}

	return db, nil
}

// createTables executes DDL migrations for projects, skills, contacts, and indexes.
func createTables(ctx context.Context, db *sql.DB) error {
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
	CREATE INDEX IF NOT EXISTS idx_projects_category ON projects(category);

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
	_, err := db.ExecContext(ctx, schema)
	return err
}

// seedInitialData inserts default portfolio projects and skills inside a database transaction.
func seedInitialData(ctx context.Context, db *sql.DB) error {
	// 1. Seed Projects if empty
	var projectCount int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM projects").Scan(&projectCount); err != nil {
		return err
	}

	if projectCount == 0 {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("failed to begin projects seed transaction: %w", err)
		}
		defer tx.Rollback()

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
				Image:        "/images/projects/safeguarding.png",
				GithubURL:    "https://github.com/VALENTINE-it/safeguarding-project",
				LiveURL:      "https://safeguarding-app-1.onrender.com/",
			},
			{
				Title:        "Nemo",
				Description:  "An award-winning IoT security system recognized at the Kijani Space Hackathon, engineered to secure, monitor, and protect fish cages in Lake Victoria through real-time telemetry and intrusion alerts.",
				Category:     "IoT & Embedded Security",
				Year:         2026,
				Technologies: `["IoT", "Embedded Systems", "Sensors", "Telemetry", "Hardware Security", "Go"]`,
				Image:        "/images/projects/fish-cage-iot.jpeg",
				GithubURL:    "https://github.com/nyabokegrace/nemo",
				LiveURL:      "",
			},
			{
				Title:        "Anga Guard",
				Description:  "A decentralized dMRV oracle and SME ESG platform empowering Western Kenya smallholders with verifiable biochar carbon removal credits compliant with Kenya National Carbon Registry.",
				Category:     "Web3 & Climate Tech Oracle",
				Year:         2026,
				Technologies: `["dMRV Oracle", "Web3", "React", "Go", "IoT Telemetry", "Smart Contracts", "ESG Analytics"]`,
				Image:        "/images/projects/angaguard.png",
				GithubURL:    "https://github.com/ClayMichael2004/angaguard",
				LiveURL:      "https://angaguard-d96o.onrender.com/",
			},
		}

		stmt, err := tx.PrepareContext(ctx, `
			INSERT INTO projects (title, description, category, year, technologies, image, github_url, live_url)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, p := range projects {
			if _, err := stmt.ExecContext(ctx, p.Title, p.Description, p.Category, p.Year, p.Technologies, p.Image, p.GithubURL, p.LiveURL); err != nil {
				return err
			}
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit projects seed transaction: %w", err)
		}
		slog.Info("Database: Successfully seeded initial portfolio projects")
	}

	// 2. Seed Skills if empty
	var skillCount int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM skills").Scan(&skillCount); err != nil {
		return err
	}

	if skillCount == 0 {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("failed to begin skills seed transaction: %w", err)
		}
		defer tx.Rollback()

		skills := []struct {
			Name     string
			Category string
		}{
			// Blockchain
			{"Smart Contracts", "Blockchain Technology Developer"},
			{"Distributed Ledgers", "Blockchain Technology Developer"},
			{"Web3 Architecture", "Blockchain Technology Developer"},
			{"Cryptographic Protocols", "Blockchain Technology Developer"},

			// Frontend
			{"HTML5", "Frontend Development"},
			{"CSS3", "Frontend Development"},
			{"JavaScript", "Frontend Development"},
			{"React", "Frontend Development"},
			{"Vite", "Frontend Development"},
			{"Responsive Design", "Frontend Development"},

			// Backend
			{"Go", "Backend Systems"},
			{"REST APIs", "Backend Systems"},
			{"HTTP Handlers", "Backend Systems"},
			{"Routing", "Backend Systems"},
			{"Concurrency", "Backend Systems"},

			// Networking & Infrastructure
			{"Networking", "Networking & Infrastructure"},
			{"TCP/IP", "Networking & Infrastructure"},
			{"Routing & Switching", "Networking & Infrastructure"},
			{"Network Security", "Networking & Infrastructure"},
			{"Protocols", "Networking & Infrastructure"},

			// Databases
			{"SQLite", "Databases & Storage"},
			{"SQL", "Databases & Storage"},
			{"Schema Design", "Databases & Storage"},
			{"Data Normalization", "Databases & Storage"},

			// Tools & DevOps
			{"Linux", "Tools & DevOps"},
			{"Git", "Tools & DevOps"},
			{"GitHub", "Tools & DevOps"},
			{"VS Code", "Tools & DevOps"},
			{"Debugging", "Tools & DevOps"},
		}

		stmt, err := tx.PrepareContext(ctx, "INSERT INTO skills (name, category) VALUES (?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, s := range skills {
			if _, err := stmt.ExecContext(ctx, s.Name, s.Category); err != nil {
				return err
			}
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit skills seed transaction: %w", err)
		}
		slog.Info("Database: Successfully seeded initial technical skills")
	}

	return nil
}
