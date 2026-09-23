package repositories

import (
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

// CreateContact persists a new contact submission into the database.
func (r *ContactRepository) CreateContact(c *models.Contact) (*models.Contact, error) {
	query := `
		INSERT INTO contacts (name, email, subject, message, created_at)
		VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)
	`
	res, err := r.db.Exec(query, c.Name, c.Email, c.Subject, c.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to insert contact: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert id: %w", err)
	}

	c.ID = id
	c.CreatedAt = time.Now().UTC()

	return c, nil
}
