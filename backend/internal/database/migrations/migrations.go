package migrations

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
)

type Migration struct {
	Version int
	Name    string
	SQL     string
}

// AllMigrations contains the ordered sequential database migrations
var AllMigrations = []Migration{
	{
		Version: 1,
		Name:    "initial_schema",
		SQL: `
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
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
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
			status TEXT NOT NULL DEFAULT 'unread',
			ip_address TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE INDEX IF NOT EXISTS idx_contacts_created_at ON contacts(created_at DESC);
		`,
	},
	{
		Version: 2,
		Name:    "admin_and_sessions",
		SQL: `
		CREATE TABLE IF NOT EXISTS admins (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE INDEX IF NOT EXISTS idx_admins_username ON admins(username);

		CREATE TABLE IF NOT EXISTS sessions (
			id TEXT PRIMARY KEY,
			admin_id INTEGER NOT NULL REFERENCES admins(id) ON DELETE CASCADE,
			expires_at DATETIME NOT NULL,
			created_at DATETIME NOT NULL,
			user_agent TEXT,
			ip_address TEXT
		);

		CREATE INDEX IF NOT EXISTS idx_sessions_expires_at ON sessions(expires_at);
		CREATE INDEX IF NOT EXISTS idx_sessions_admin_id ON sessions(admin_id);
		`,
	},
	{
		Version: 3,
		Name:    "contacts_status_index",
		SQL: `
		CREATE INDEX IF NOT EXISTS idx_contacts_status ON contacts(status);
		`,
	},
}

// Run executes all pending migrations inside transactions
func Run(db *sql.DB) error {
	// Create schema_migrations table if not exists
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// Upgrade legacy schema if upgrading existing database
	if err := upgradeLegacyColumns(db); err != nil {
		slog.Warn("Legacy column upgrade notice", "error", err)
	}

	for _, m := range AllMigrations {
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = ?)", m.Version).Scan(&exists)
		if err != nil {
			return fmt.Errorf("failed to query migration status for v%d: %w", m.Version, err)
		}

		if exists {
			continue
		}

		slog.Info("Applying database migration", "version", m.Version, "name", m.Name)

		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction for migration v%d: %w", m.Version, err)
		}

		if _, err := tx.Exec(m.SQL); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("migration v%d (%s) failed: %w", m.Version, m.Name, err)
		}

		if _, err := tx.Exec("INSERT INTO schema_migrations (version, name) VALUES (?, ?)", m.Version, m.Name); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("failed to record migration v%d: %w", m.Version, err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit migration v%d: %w", m.Version, err)
		}

		slog.Info("Applied database migration successfully", "version", m.Version, "name", m.Name)
	}

	return nil
}

func upgradeLegacyColumns(db *sql.DB) error {
	// Ensure contacts has status and ip_address
	_ = addColumnIfNotExists(db, "contacts", "status", "TEXT NOT NULL DEFAULT 'unread'")
	_ = addColumnIfNotExists(db, "contacts", "ip_address", "TEXT")
	// Ensure projects has updated_at (SQLite does not allow CURRENT_TIMESTAMP in ALTER TABLE ADD COLUMN)
	if err := addColumnIfNotExists(db, "projects", "updated_at", "DATETIME"); err == nil {
		_, _ = db.Exec("UPDATE projects SET updated_at = created_at WHERE updated_at IS NULL")
	}
	return nil
}

func addColumnIfNotExists(db *sql.DB, table, column, colDef string) error {
	rows, err := db.Query(fmt.Sprintf("PRAGMA table_info(%s);", table))
	if err != nil {
		return err
	}
	defer rows.Close()

	exists := false
	for rows.Next() {
		var cid int
		var name, ctype string
		var notnull, pk int
		var dfltValue sql.NullString
		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk); err == nil {
			if strings.EqualFold(name, column) {
				exists = true
				break
			}
		}
	}

	if !exists {
		alterQuery := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", table, column, colDef)
		if _, err := db.Exec(alterQuery); err != nil {
			return err
		}
		slog.Info("Upgraded legacy table schema: added column", "table", table, "column", column)
	}

	return nil
}
