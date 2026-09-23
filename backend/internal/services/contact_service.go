package services

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var msgs []string
	for _, err := range v {
		msgs = append(msgs, fmt.Sprintf("%s: %s", err.Field, err.Message))
	}
	return strings.Join(msgs, "; ")
}

type ContactService struct {
	repo *repositories.ContactRepository
}

func NewContactService(repo *repositories.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

// ValidateContactRequest checks all fields according to security and business rules.
func (s *ContactService) ValidateContactRequest(req *models.ContactRequest) ValidationErrors {
	var errs ValidationErrors

	name := strings.TrimSpace(req.Name)
	if name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "Name is required."})
	} else if len(name) > 100 {
		errs = append(errs, ValidationError{Field: "name", Message: "Name cannot exceed 100 characters."})
	}

	email := strings.TrimSpace(req.Email)
	if email == "" {
		errs = append(errs, ValidationError{Field: "email", Message: "Email is required."})
	} else if len(email) < 3 || len(email) > 254 {
		errs = append(errs, ValidationError{Field: "email", Message: "Email must be between 3 and 254 characters."})
	} else {
		// Verify email structure
		addr, err := mail.ParseAddress(email)
		if err != nil || addr.Address != email || !strings.Contains(email, ".") {
			errs = append(errs, ValidationError{Field: "email", Message: "Please provide a valid email address."})
		}
	}

	subject := strings.TrimSpace(req.Subject)
	if subject == "" {
		errs = append(errs, ValidationError{Field: "subject", Message: "Subject is required."})
	} else if len(subject) > 200 {
		errs = append(errs, ValidationError{Field: "subject", Message: "Subject cannot exceed 200 characters."})
	}

	message := strings.TrimSpace(req.Message)
	if message == "" {
		errs = append(errs, ValidationError{Field: "message", Message: "Message is required."})
	} else if len(message) < 5 {
		errs = append(errs, ValidationError{Field: "message", Message: "Message must be at least 5 characters long."})
	} else if len(message) > 5000 {
		errs = append(errs, ValidationError{Field: "message", Message: "Message cannot exceed 5000 characters."})
	}

	return errs
}

// SubmitContact validates and stores a contact message.
func (s *ContactService) SubmitContact(req *models.ContactRequest) (*models.Contact, error) {
	if req == nil {
		return nil, errors.New("contact request cannot be nil")
	}

	if errs := s.ValidateContactRequest(req); len(errs) > 0 {
		return nil, errs
	}

	contact := &models.Contact{
		Name:    strings.TrimSpace(req.Name),
		Email:   strings.TrimSpace(req.Email),
		Subject: strings.TrimSpace(req.Subject),
		Message: strings.TrimSpace(req.Message),
	}

	saved, err := s.repo.CreateContact(contact)
	if err != nil {
		return nil, fmt.Errorf("contact_service: failed to save message: %w", err)
	}

	return saved, nil
}
