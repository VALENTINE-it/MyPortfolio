package repositories

import (
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

// GetAllProjects retrieves all projects ordered by year descending, then id descending.
func (r *ProjectRepository) GetAllProjects() ([]models.Project, error) {
	query := `
		SELECT id, title, description, category, year, technologies, image, github_url, live_url, created_at
		FROM projects
		ORDER BY year DESC, id ASC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %w", err)
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		var techJSON string
		var githubURL, liveURL sql.NullString
		var createdAtStr sql.NullString

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
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}

		if githubURL.Valid {
			p.GithubURL = githubURL.String
		}
		if liveURL.Valid {
			p.LiveURL = liveURL.String
		}

		// Unmarshal technologies JSON
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

		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating project rows: %w", err)
	}

	if projects == nil {
		projects = []models.Project{}
	}

	return projects, nil
}

// GetProjectByID retrieves a single project by its ID.
func (r *ProjectRepository) GetProjectByID(id int64) (*models.Project, error) {
	query := `
		SELECT id, title, description, category, year, technologies, image, github_url, live_url, created_at
		FROM projects
		WHERE id = ?
	`
	row := r.db.QueryRow(query, id)

	var p models.Project
	var techJSON string
	var githubURL, liveURL sql.NullString
	var createdAtStr sql.NullString

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
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProjectNotFound
		}
		return nil, fmt.Errorf("failed to scan project by id: %w", err)
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

	return &p, nil
}
