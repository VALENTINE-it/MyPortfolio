package models

import "time"

// Contact represents a stored message submitted through the contact form.
type Contact struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// ContactRequest represents the incoming payload from the frontend contact form.
type ContactRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Subject  string `json:"subject"`
	Message  string `json:"message"`
	Honeypot string `json:"honeypot,omitempty"` // Anti-spam trap: bots fill hidden fields, humans do not
}
