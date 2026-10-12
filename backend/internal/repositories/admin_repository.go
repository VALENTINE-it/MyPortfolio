package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"portfolio-backend/internal/models"
)

var (
	ErrAdminNotFound = errors.New("administrator not found")
)

type AdminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

// GetByUsername retrieves an admin record by username
func (r *AdminRepository) GetByUsername(ctx context.Context, username string) (*models.Admin, error) {
	query := `
		SELECT id, username, password_hash, created_at, updated_at
		FROM admins
		WHERE username = ?
	`
	row := r.db.QueryRowContext(ctx, query, username)

	var a models.Admin
	var createdAtStr, updatedAtStr sql.NullString
	if err := row.Scan(&a.ID, &a.Username, &a.PasswordHash, &createdAtStr, &updatedAtStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAdminNotFound
		}
		return nil, fmt.Errorf("failed to get admin by username: %w", err)
	}

	if createdAtStr.Valid {
		if t, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
			a.CreatedAt = t
		}
	}
	if updatedAtStr.Valid {
		if t, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
			a.UpdatedAt = t
		}
	}

	return &a, nil
}

// GetByID retrieves an admin record by ID
func (r *AdminRepository) GetByID(ctx context.Context, id int64) (*models.Admin, error) {
	query := `
		SELECT id, username, password_hash, created_at, updated_at
		FROM admins
		WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var a models.Admin
	var createdAtStr, updatedAtStr sql.NullString
	if err := row.Scan(&a.ID, &a.Username, &a.PasswordHash, &createdAtStr, &updatedAtStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAdminNotFound
		}
		return nil, fmt.Errorf("failed to get admin by id: %w", err)
	}

	if createdAtStr.Valid {
		if t, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
			a.CreatedAt = t
		}
	}
	if updatedAtStr.Valid {
		if t, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
			a.UpdatedAt = t
		}
	}

	return &a, nil
}

// Create persists a new administrator
func (r *AdminRepository) Create(ctx context.Context, username, passwordHash string) (*models.Admin, error) {
	query := `
		INSERT INTO admins (username, password_hash, created_at, updated_at)
		VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	res, err := r.db.ExecContext(ctx, query, username, passwordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to create admin: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get admin id: %w", err)
	}

	return &models.Admin{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}, nil
}

// Count returns the total number of registered administrators
func (r *AdminRepository) Count(ctx context.Context) (int, error) {
	var count int
	if err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM admins").Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
