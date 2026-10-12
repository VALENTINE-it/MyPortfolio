package models

import "time"

// Skill represents a technical skill categorized by domain.
type Skill struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// SkillInput represents payload for creating or updating a skill.
type SkillInput struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

// SkillsByCategory groups skills under their respective categories.
type SkillsByCategory struct {
	Category string   `json:"category"`
	Skills   []string `json:"skills"`
}
