package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"portfolio-backend/internal/models"
)

type SkillRepository struct {
	db *sql.DB
}

func NewSkillRepository(db *sql.DB) *SkillRepository {
	return &SkillRepository{db: db}
}

// GetAllSkills retrieves all skills, optionally filtered by category.
func (r *SkillRepository) GetAllSkills(category string) ([]models.Skill, error) {
	var rows *sql.Rows
	var err error

	if category != "" {
		query := `SELECT id, name, category, created_at FROM skills WHERE category = ? ORDER BY id ASC`
		rows, err = r.db.Query(query, category)
	} else {
		query := `SELECT id, name, category, created_at FROM skills ORDER BY id ASC`
		rows, err = r.db.Query(query)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query skills: %w", err)
	}
	defer rows.Close()

	var skills []models.Skill
	for rows.Next() {
		var s models.Skill
		var createdAtStr sql.NullString

		if err := rows.Scan(&s.ID, &s.Name, &s.Category, &createdAtStr); err != nil {
			return nil, fmt.Errorf("failed to scan skill: %w", err)
		}

		if createdAtStr.Valid {
			if t, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				s.CreatedAt = t
			}
		}

		skills = append(skills, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating skill rows: %w", err)
	}

	if skills == nil {
		skills = []models.Skill{}
	}

	return skills, nil
}
