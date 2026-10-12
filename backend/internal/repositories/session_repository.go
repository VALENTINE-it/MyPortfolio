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
	ErrSessionNotFound = errors.New("session not found")
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

// CreateSession persists a new session token
func (r *SessionRepository) CreateSession(ctx context.Context, s *models.Session) error {
	query := `
		INSERT INTO sessions (id, admin_id, expires_at, created_at, user_agent, ip_address)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	expiresAtStr := s.ExpiresAt.UTC().Format(time.RFC3339)
	createdAtStr := s.CreatedAt.UTC().Format(time.RFC3339)

	_, err := r.db.ExecContext(ctx, query, s.ID, s.AdminID, expiresAtStr, createdAtStr, s.UserAgent, s.IPAddress)
	if err != nil {
		return fmt.Errorf("failed to insert session: %w", err)
	}
	return nil
}

// GetSession retrieves an active session by opaque session ID
func (r *SessionRepository) GetSession(ctx context.Context, id string) (*models.Session, error) {
	query := `
		SELECT id, admin_id, expires_at, created_at, user_agent, ip_address
		FROM sessions
		WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var s models.Session
	var expiresAtStr, createdAtStr string
	var ua, ip sql.NullString

	if err := row.Scan(&s.ID, &s.AdminID, &expiresAtStr, &createdAtStr, &ua, &ip); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSessionNotFound
		}
		return nil, fmt.Errorf("failed to query session: %w", err)
	}

	if ua.Valid {
		s.UserAgent = ua.String
	}
	if ip.Valid {
		s.IPAddress = ip.String
	}

	s.ExpiresAt = parseFlexibleTime(expiresAtStr)
	s.CreatedAt = parseFlexibleTime(createdAtStr)

	return &s, nil
}

// DeleteSession destroys a session upon logout
func (r *SessionRepository) DeleteSession(ctx context.Context, id string) error {
	query := `DELETE FROM sessions WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// DeleteExpiredSessions prunes sessions that have passed their expiration timestamp
func (r *SessionRepository) DeleteExpiredSessions(ctx context.Context) (int64, error) {
	nowStr := time.Now().UTC().Format(time.RFC3339)
	query := `DELETE FROM sessions WHERE expires_at < ?`
	res, err := r.db.ExecContext(ctx, query, nowStr)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteAdminSessions invalidates all sessions for a specific administrator
func (r *SessionRepository) DeleteAdminSessions(ctx context.Context, adminID int64) error {
	query := `DELETE FROM sessions WHERE admin_id = ?`
	_, err := r.db.ExecContext(ctx, query, adminID)
	return err
}

func parseFlexibleTime(str string) time.Time {
	formats := []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
	}
	for _, f := range formats {
		if t, err := time.Parse(f, str); err == nil {
			return t.UTC()
		}
	}
	return time.Time{}
}
