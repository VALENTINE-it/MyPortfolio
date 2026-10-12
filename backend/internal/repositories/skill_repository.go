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
	ErrSkillNotFound = errors.New("skill not found")
)

type SkillRepository struct {
	db *sql.DB
}

func NewSkillRepository(db *sql.DB) *SkillRepository {
	return &SkillRepository{db: db}
}

// GetAllSkills retrieves all skills, optionally filtered by category using parameterized query.
func (r *SkillRepository) GetAllSkills(ctx context.Context, category string) ([]models.Skill, error) {
	var rows *sql.Rows
	var err error

	if category != "" {
		query := `SELECT id, name, category, created_at FROM skills WHERE category = ? ORDER BY id ASC`
		rows, err = r.db.QueryContext(ctx, query, category)
	} else {
		query := `SELECT id, name, category, created_at FROM skills ORDER BY id ASC`
		rows, err = r.db.QueryContext(ctx, query)
	}

	if err != nil {
		return nil, fmt.Errorf("skill_repo: failed to query skills: %w", err)
	}
	defer rows.Close()

	var skills []models.Skill
	for rows.Next() {
		var s models.Skill
		var createdAtStr sql.NullString

		if err := rows.Scan(&s.ID, &s.Name, &s.Category, &createdAtStr); err != nil {
			return nil, fmt.Errorf("skill_repo: failed to scan skill: %w", err)
		}

		if createdAtStr.Valid {
			if t, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				s.CreatedAt = t
			}
		}

		skills = append(skills, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("skill_repo: error iterating skill rows: %w", err)
	}

	if skills == nil {
		skills = []models.Skill{}
	}

	return skills, nil
}

// GetByID retrieves a single skill by ID
func (r *SkillRepository) GetByID(ctx context.Context, id int64) (*models.Skill, error) {
	query := `SELECT id, name, category, created_at FROM skills WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var s models.Skill
	var createdAtStr sql.NullString
	if err := row.Scan(&s.ID, &s.Name, &s.Category, &createdAtStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSkillNotFound
		}
		return nil, fmt.Errorf("failed to get skill by id: %w", err)
	}

	if createdAtStr.Valid {
		if t, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
			s.CreatedAt = t
		}
	}
	return &s, nil
}

// Create inserts a new skill
func (r *SkillRepository) Create(ctx context.Context, skill *models.Skill) (*models.Skill, error) {
	query := `INSERT INTO skills (name, category, created_at) VALUES (?, ?, CURRENT_TIMESTAMP)`
	res, err := r.db.ExecContext(ctx, query, skill.Name, skill.Category)
	if err != nil {
		return nil, fmt.Errorf("failed to insert skill: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert id: %w", err)
	}

	skill.ID = id
	skill.CreatedAt = time.Now().UTC()
	return skill, nil
}

// Update modifies an existing skill
func (r *SkillRepository) Update(ctx context.Context, skill *models.Skill) error {
	query := `UPDATE skills SET name = ?, category = ? WHERE id = ?`
	res, err := r.db.ExecContext(ctx, query, skill.Name, skill.Category, skill.ID)
	if err != nil {
		return fmt.Errorf("failed to update skill: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrSkillNotFound
	}
	return nil
}

// Delete removes a skill by ID
func (r *SkillRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM skills WHERE id = ?`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete skill: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrSkillNotFound
	}
	return nil
}
