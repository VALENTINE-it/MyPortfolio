package services

import (
	"context"
	"fmt"
	"strings"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/utils"
)

type SkillService struct {
	repo *repositories.SkillRepository
}

func NewSkillService(repo *repositories.SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

// GetAllSkills retrieves all skills optionally filtered by category.
func (s *SkillService) GetAllSkills(ctx context.Context, category string) ([]models.Skill, error) {
	skills, err := s.repo.GetAllSkills(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("skill_service: %w", err)
	}
	return skills, nil
}

// GetSkillByID retrieves a skill by ID
func (s *SkillService) GetSkillByID(ctx context.Context, id int64) (*models.Skill, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid skill ID: must be positive")
	}
	return s.repo.GetByID(ctx, id)
}

// CreateSkill validates input and inserts a new skill
func (s *SkillService) CreateSkill(ctx context.Context, input *models.SkillInput) (*models.Skill, error) {
	name := utils.SanitizeString(input.Name)
	category := utils.SanitizeString(input.Category)

	if name == "" || len(name) > 100 {
		return nil, fmt.Errorf("skill name is required and must not exceed 100 characters")
	}
	if category == "" || len(category) > 100 {
		return nil, fmt.Errorf("skill category is required and must not exceed 100 characters")
	}

	skill := &models.Skill{
		Name:     name,
		Category: category,
	}

	return s.repo.Create(ctx, skill)
}

// UpdateSkill updates an existing skill
func (s *SkillService) UpdateSkill(ctx context.Context, id int64, input *models.SkillInput) (*models.Skill, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid skill ID: must be positive")
	}

	name := utils.SanitizeString(input.Name)
	category := utils.SanitizeString(input.Category)

	if name == "" || len(name) > 100 {
		return nil, fmt.Errorf("skill name is required and must not exceed 100 characters")
	}
	if category == "" || len(category) > 100 {
		return nil, fmt.Errorf("skill category is required and must not exceed 100 characters")
	}

	skill := &models.Skill{
		ID:       id,
		Name:     name,
		Category: category,
	}

	if err := s.repo.Update(ctx, skill); err != nil {
		return nil, err
	}
	return skill, nil
}

// DeleteSkill deletes a skill by ID
func (s *SkillService) DeleteSkill(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid skill ID: must be positive")
	}
	return s.repo.Delete(ctx, id)
}

// GetSkillsGrouped returns skills grouped by their categories in defined order.
func (s *SkillService) GetSkillsGrouped(ctx context.Context) ([]models.SkillsByCategory, error) {
	skills, err := s.repo.GetAllSkills(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("skill_service: %w", err)
	}

	categoryOrder := []string{
		"Frontend",
		"Backend",
		"Tools",
		"Blockchain Technology Developer",
		"Other",
	}
	groupedMap := make(map[string][]string)

	for _, sk := range skills {
		groupedMap[sk.Category] = append(groupedMap[sk.Category], sk.Name)
	}

	var result []models.SkillsByCategory
	seen := make(map[string]bool)

	// Add known categories in prioritized order
	for _, cat := range categoryOrder {
		if items, exists := groupedMap[cat]; exists && len(items) > 0 {
			result = append(result, models.SkillsByCategory{
				Category: cat,
				Skills:   items,
			})
			seen[cat] = true
		}
	}

	// Add any additional unexpected categories
	for cat, items := range groupedMap {
		if !seen[cat] && len(items) > 0 {
			result = append(result, models.SkillsByCategory{
				Category: cat,
				Skills:   items,
			})
		}
	}

	if result == nil {
		result = []models.SkillsByCategory{}
	}

	return result, nil
}

func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
