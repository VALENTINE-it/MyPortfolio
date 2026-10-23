package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/utils"
)

var (
	ErrInvalidProjectID = errors.New("project id must be greater than zero")
)

type ProjectService struct {
	repo *repositories.ProjectRepository
}

func NewProjectService(repo *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

// GetAllProjects retrieves all projects from repository.
func (s *ProjectService) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	projects, err := s.repo.GetAllProjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("project_service: %w", err)
	}
	return projects, nil
}

// GetProjectByID validates ID and retrieves project by ID.
func (s *ProjectService) GetProjectByID(ctx context.Context, id int64) (*models.Project, error) {
	if id <= 0 {
		return nil, ErrInvalidProjectID
	}
	project, err := s.repo.GetProjectByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// CreateProject validates input and inserts a new project
func (s *ProjectService) CreateProject(ctx context.Context, input *models.ProjectInput) (*models.Project, error) {
	if err := s.validateProjectInput(input); err != nil {
		return nil, err
	}

	p := &models.Project{
		Title:        utils.SanitizeString(input.Title),
		Description:  utils.SanitizeString(input.Description),
		Category:     utils.SanitizeString(input.Category),
		Year:         input.Year,
		Technologies: input.Technologies,
		Image:        utils.CleanPlainString(input.Image),
		GithubURL:    utils.CleanPlainString(input.GithubURL),
		LiveURL:      utils.CleanPlainString(input.LiveURL),
	}

	if p.Technologies == nil {
		p.Technologies = []string{}
	}

	return s.repo.Create(ctx, p)
}

// UpdateProject validates input and updates an existing project
func (s *ProjectService) UpdateProject(ctx context.Context, id int64, input *models.ProjectInput) (*models.Project, error) {
	if id <= 0 {
		return nil, ErrInvalidProjectID
	}
	if err := s.validateProjectInput(input); err != nil {
		return nil, err
	}

	existing, err := s.repo.GetProjectByID(ctx, id)
	if err != nil {
		return nil, err
	}

	existing.Title = utils.SanitizeString(input.Title)
	existing.Description = utils.SanitizeString(input.Description)
	existing.Category = utils.SanitizeString(input.Category)
	existing.Year = input.Year
	existing.Technologies = input.Technologies
	if existing.Technologies == nil {
		existing.Technologies = []string{}
	}
	existing.Image = utils.CleanPlainString(input.Image)
	existing.GithubURL = utils.CleanPlainString(input.GithubURL)
	existing.LiveURL = utils.CleanPlainString(input.LiveURL)

	if err := s.repo.Update(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

// DeleteProject removes a project
func (s *ProjectService) DeleteProject(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrInvalidProjectID
	}
	return s.repo.Delete(ctx, id)
}

func (s *ProjectService) validateProjectInput(in *models.ProjectInput) error {
	if in.Title == "" || len(in.Title) > 200 {
		return fmt.Errorf("title is required and must not exceed 200 characters")
	}
	if in.Description == "" || len(in.Description) > 5000 {
		return fmt.Errorf("description is required and must not exceed 5000 characters")
	}
	if in.Category == "" || len(in.Category) > 100 {
		return fmt.Errorf("category is required and must not exceed 100 characters")
	}
	currentYear := time.Now().Year()
	if in.Year < 2000 || in.Year > currentYear+5 {
		return fmt.Errorf("year must be between 2000 and %d", currentYear+5)
	}
	if in.Image == "" || len(in.Image) > 500 {
		return fmt.Errorf("image path is required and must not exceed 500 characters")
	}
	return nil
}

// GetProjectsPaginated validates pagination arguments and retrieves paginated projects.
func (s *ProjectService) GetProjectsPaginated(ctx context.Context, page, limit int) ([]models.Project, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 3
	}
	if limit > 50 {
		limit = 50
	}
	offset := (page - 1) * limit
	return s.repo.GetProjectsPaginated(ctx, limit, offset)
}
