package services

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/utils"
)

var (
	ErrSpamDetected = errors.New("spam detected")
	emailRegex      = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
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

// ValidateContactRequest checks all fields according to strict security and length rules.
func (s *ContactService) ValidateContactRequest(req *models.ContactRequest) ValidationErrors {
	var errs ValidationErrors

	// Honeypot anti-spam check: if bot filled this hidden field, fail validation
	if strings.TrimSpace(req.Honeypot) != "" {
		errs = append(errs, ValidationError{Field: "honeypot", Message: "Spam submission detected."})
		return errs
	}

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
		// Strict address parsing and regex verification
		addr, err := mail.ParseAddress(email)
		if err != nil || addr.Address != email || !emailRegex.MatchString(email) {
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

// SubmitContact validates, sanitizes, and stores a contact message.
func (s *ContactService) SubmitContact(ctx context.Context, req *models.ContactRequest) (*models.Contact, error) {
	if req == nil {
		return nil, errors.New("contact request cannot be nil")
	}

	if errs := s.ValidateContactRequest(req); len(errs) > 0 {
		return nil, errs
	}

	contact := &models.Contact{
		Name:    utils.SanitizeString(req.Name),
		Email:   utils.CleanPlainString(req.Email),
		Subject: utils.SanitizeString(req.Subject),
		Message: utils.SanitizeString(req.Message),
		Status:  "unread",
	}

	saved, err := s.repo.CreateContact(ctx, contact)
	if err != nil {
		return nil, fmt.Errorf("contact_service: failed to save message: %w", err)
	}

	return saved, nil
}

// GetMessages retrieves paginated contact messages for administrative review
func (s *ContactService) GetMessages(ctx context.Context, limit, offset int, status string) ([]models.Contact, int, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.GetAllMessages(ctx, limit, offset, status)
}

// GetMessageByID retrieves a single message by ID
func (s *ContactService) GetMessageByID(ctx context.Context, id int64) (*models.Contact, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid message ID")
	}
	return s.repo.GetByID(ctx, id)
}

// UpdateMessageStatus updates the status of a message (e.g. "read", "archived")
func (s *ContactService) UpdateMessageStatus(ctx context.Context, id int64, status string) error {
	if id <= 0 {
		return fmt.Errorf("invalid message ID")
	}
	validStatuses := map[string]bool{
		"unread":   true,
		"read":     true,
		"archived": true,
	}
	if !validStatuses[status] {
		return fmt.Errorf("invalid status: must be unread, read, or archived")
	}
	return s.repo.UpdateStatus(ctx, id, status)
}

// DeleteMessage deletes a message
func (s *ContactService) DeleteMessage(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid message ID")
	}
	return s.repo.Delete(ctx, id)
}
