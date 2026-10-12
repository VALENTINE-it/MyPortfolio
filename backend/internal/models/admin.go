package models

import "time"

// Admin represents an authorized system administrator
type Admin struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // Never serialized in JSON
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// AdminLoginRequest represents the login credentials payload
type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Session represents an active authenticated administrator session
type Session struct {
	ID        string    `json:"id"`
	AdminID   int64     `json:"adminId"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	UserAgent string    `json:"userAgent"`
	IPAddress string    `json:"ipAddress"`
}

// AdminProfile represents public admin profile details
type AdminProfile struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
