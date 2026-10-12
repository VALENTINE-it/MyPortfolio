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
	ErrContactNotFound = errors.New("contact message not found")
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

// CreateContact persists a new contact submission into SQLite using ExecContext with parameterized query.
func (r *ContactRepository) CreateContact(ctx context.Context, c *models.Contact) (*models.Contact, error) {
	if c.Status == "" {
		c.Status = "unread"
	}
	query := `
		INSERT INTO contacts (name, email, subject, message, status, ip_address, created_at)
		VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
	`
	res, err := r.db.ExecContext(ctx, query, c.Name, c.Email, c.Subject, c.Message, c.Status, c.IPAddress)
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

// GetAllMessages retrieves paginated messages with optional status filtering for admin review
func (r *ContactRepository) GetAllMessages(ctx context.Context, limit, offset int, status string) ([]models.Contact, int, error) {
	var total int
	countQuery := `SELECT COUNT(*) FROM contacts`
	var rows *sql.Rows
	var err error

	if status != "" {
		countQuery += ` WHERE status = ?`
		if err := r.db.QueryRowContext(ctx, countQuery, status).Scan(&total); err != nil {
			return nil, 0, err
		}

		query := `
			SELECT id, name, email, subject, message, status, ip_address, created_at
			FROM contacts
			WHERE status = ?
			ORDER BY created_at DESC
			LIMIT ? OFFSET ?
		`
		rows, err = r.db.QueryContext(ctx, query, status, limit, offset)
	} else {
		if err := r.db.QueryRowContext(ctx, countQuery).Scan(&total); err != nil {
			return nil, 0, err
		}

		query := `
			SELECT id, name, email, subject, message, status, ip_address, created_at
			FROM contacts
			ORDER BY created_at DESC
			LIMIT ? OFFSET ?
		`
		rows, err = r.db.QueryContext(ctx, query, limit, offset)
	}

	if err != nil {
		return nil, 0, fmt.Errorf("failed to query contact messages: %w", err)
	}
	defer rows.Close()

	var list []models.Contact
	for rows.Next() {
		var c models.Contact
		var ipStr sql.NullString
		var createdAtStr sql.NullString

		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Subject, &c.Message, &c.Status, &ipStr, &createdAtStr); err != nil {
			return nil, 0, fmt.Errorf("failed to scan contact: %w", err)
		}
		if ipStr.Valid {
			c.IPAddress = ipStr.String
		}
		if createdAtStr.Valid {
			if t, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				c.CreatedAt = t
			}
		}
		list = append(list, c)
	}

	if list == nil {
		list = []models.Contact{}
	}

	return list, total, nil
}

// GetByID retrieves a single message by ID
func (r *ContactRepository) GetByID(ctx context.Context, id int64) (*models.Contact, error) {
	query := `SELECT id, name, email, subject, message, status, ip_address, created_at FROM contacts WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var c models.Contact
	var ipStr, createdAtStr sql.NullString
	if err := row.Scan(&c.ID, &c.Name, &c.Email, &c.Subject, &c.Message, &c.Status, &ipStr, &createdAtStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrContactNotFound
		}
		return nil, fmt.Errorf("failed to retrieve contact message: %w", err)
	}

	if ipStr.Valid {
		c.IPAddress = ipStr.String
	}
	if createdAtStr.Valid {
		if t, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
			c.CreatedAt = t
		}
	}

	return &c, nil
}

// UpdateStatus changes message status (e.g. "read", "unread", "archived")
func (r *ContactRepository) UpdateStatus(ctx context.Context, id int64, status string) error {
	query := `UPDATE contacts SET status = ? WHERE id = ?`
	res, err := r.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("failed to update contact status: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrContactNotFound
	}

	return nil
}

// Delete removes a contact message by ID
func (r *ContactRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM contacts WHERE id = ?`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete contact: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrContactNotFound
	}

	return nil
}
