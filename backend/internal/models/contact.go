package models

import "time"

// Contact represents a message submitted through the contact form.
type Contact struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// ContactRequest represents the payload incoming from the frontend contact form.
type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}
