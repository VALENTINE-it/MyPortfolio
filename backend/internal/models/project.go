package models

import "time"

// Project represents a portfolio project item.
type Project struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	Year         int       `json:"year"`
	Technologies []string  `json:"technologies"`
	Image        string    `json:"image"`
	GithubURL    string    `json:"githubUrl"`
	LiveURL      string    `json:"liveUrl"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}
