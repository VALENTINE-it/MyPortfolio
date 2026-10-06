package services

import (
	"context"
	"fmt"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
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

// GetSkillsGrouped returns skills grouped by their categories in defined order.
func (s *SkillService) GetSkillsGrouped(ctx context.Context) ([]models.SkillsByCategory, error) {
	skills, err := s.repo.GetAllSkills(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("skill_service: %w", err)
	}

	categoryOrder := []string{
		"Blockchain Technology Developer",
		"Frontend Development",
		"Backend Systems",
		"Networking & Infrastructure",
		"Databases & Storage",
		"Tools & DevOps",
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
