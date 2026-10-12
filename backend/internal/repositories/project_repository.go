package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"portfolio-backend/internal/models"
)

var (
	ErrProjectNotFound = errors.New("project not found")
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

// GetAllProjects retrieves all projects ordered by year descending, then id ascending.
func (r *ProjectRepository) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	query := `
		SELECT id, title, description, category, year, technologies, image, github_url, live_url, created_at, updated_at
		FROM projects
		ORDER BY year DESC, id ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("project_repo: failed to query projects: %w", err)
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		var techJSON string
		var githubURL, liveURL sql.NullString
		var createdAtStr, updatedAtStr sql.NullString

		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.Category,
			&p.Year,
			&techJSON,
			&p.Image,
			&githubURL,
			&liveURL,
			&createdAtStr,
			&updatedAtStr,
		)
		if err != nil {
			return nil, fmt.Errorf("project_repo: failed to scan project: %w", err)
		}

		if githubURL.Valid {
			p.GithubURL = githubURL.String
		}
		if liveURL.Valid {
			p.LiveURL = liveURL.String
		}

		if techJSON != "" {
			if err := json.Unmarshal([]byte(techJSON), &p.Technologies); err != nil {
				p.Technologies = []string{}
			}
		} else {
			p.Technologies = []string{}
		}

		if createdAtStr.Valid {
			if t, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				p.CreatedAt = t
			}
		}
		if updatedAtStr.Valid {
			if t, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
				p.UpdatedAt = t
			}
		}

		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("project_repo: error iterating project rows: %w", err)
	}

	if projects == nil {
		projects = []models.Project{}
	}

	return projects, nil
}

// GetProjectByID retrieves a single project by its ID using parameterized query.
func (r *ProjectRepository) GetProjectByID(ctx context.Context, id int64) (*models.Project, error) {
	query := `
		SELECT id, title, description, category, year, technologies, image, github_url, live_url, created_at, updated_at
		FROM projects
		WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var p models.Project
	var techJSON string
	var githubURL, liveURL sql.NullString
	var createdAtStr, updatedAtStr sql.NullString

	err := row.Scan(
		&p.ID,
		&p.Title,
		&p.Description,
		&p.Category,
		&p.Year,
		&techJSON,
		&p.Image,
		&githubURL,
		&liveURL,
		&createdAtStr,
		&updatedAtStr,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProjectNotFound
		}
		return nil, fmt.Errorf("project_repo: failed to scan project by id: %w", err)
	}

	if githubURL.Valid {
		p.GithubURL = githubURL.String
	}
	if liveURL.Valid {
		p.LiveURL = liveURL.String
	}

	if techJSON != "" {
		if err := json.Unmarshal([]byte(techJSON), &p.Technologies); err != nil {
			p.Technologies = []string{}
		}
	} else {
		p.Technologies = []string{}
	}

	if createdAtStr.Valid {
		if t, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
			p.CreatedAt = t
		}
	}
	if updatedAtStr.Valid {
		if t, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
			p.UpdatedAt = t
		}
	}

	return &p, nil
}

// Create inserts a new project
func (r *ProjectRepository) Create(ctx context.Context, p *models.Project) (*models.Project, error) {
	techJSON, err := json.Marshal(p.Technologies)
	if err != nil {
		techJSON = []byte("[]")
	}

	query := `
		INSERT INTO projects (title, description, category, year, technologies, image, github_url, live_url, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	res, err := r.db.ExecContext(ctx, query, p.Title, p.Description, p.Category, p.Year, string(techJSON), p.Image, p.GithubURL, p.LiveURL)
	if err != nil {
		return nil, fmt.Errorf("failed to insert project: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get project insert id: %w", err)
	}

	p.ID = id
	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = p.CreatedAt
	return p, nil
}

// Update modifies an existing project
func (r *ProjectRepository) Update(ctx context.Context, p *models.Project) error {
	techJSON, err := json.Marshal(p.Technologies)
	if err != nil {
		techJSON = []byte("[]")
	}

	query := `
		UPDATE projects
		SET title = ?, description = ?, category = ?, year = ?, technologies = ?, image = ?, github_url = ?, live_url = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`
	res, err := r.db.ExecContext(ctx, query, p.Title, p.Description, p.Category, p.Year, string(techJSON), p.Image, p.GithubURL, p.LiveURL, p.ID)
	if err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrProjectNotFound
	}

	return nil
}

// Delete removes a project by ID
func (r *ProjectRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM projects WHERE id = ?`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrProjectNotFound
	}

	return nil
}
