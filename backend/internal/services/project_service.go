package services

import (
	"errors"
	"fmt"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
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
func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
	projects, err := s.repo.GetAllProjects()
	if err != nil {
		return nil, fmt.Errorf("project_service: %w", err)
	}
	return projects, nil
}

// GetProjectByID validates ID and retrieves project by ID.
func (s *ProjectService) GetProjectByID(id int64) (*models.Project, error) {
	if id <= 0 {
		return nil, ErrInvalidProjectID
	}
	project, err := s.repo.GetProjectByID(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}
