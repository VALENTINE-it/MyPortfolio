package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"portfolio-backend/internal/models"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

// CreateContact persists a new contact submission into SQLite using ExecContext with parameterized query.
func (r *ContactRepository) CreateContact(ctx context.Context, c *models.Contact) (*models.Contact, error) {
	query := `
		INSERT INTO contacts (name, email, subject, message, created_at)
		VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)
	`
	res, err := r.db.ExecContext(ctx, query, c.Name, c.Email, c.Subject, c.Message)
	if err != nil {
		return nil, fmt.Errorf("contact_repo: failed to insert contact: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("contact_repo: failed to retrieve last insert id: %w", err)
	}

	c.ID = id
	c.CreatedAt = time.Now().UTC()

	return c, nil
}
