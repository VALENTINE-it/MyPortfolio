package utils

import (
	"html"
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	// Regex for valid username (letters, digits, underscores, dashes, 3-32 chars)
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_\-\.]{3,32}$`)
)

// SanitizeString trims whitespace, removes null bytes, and escapes HTML tags to prevent XSS
func SanitizeString(s string) string {
	cleaned := strings.ReplaceAll(s, "\x00", "")
	cleaned = strings.TrimSpace(cleaned)
	return html.EscapeString(cleaned)
}

// CleanPlainString trims whitespace and removes control characters without escaping standard punctuation
func CleanPlainString(s string) string {
	cleaned := strings.ReplaceAll(s, "\x00", "")
	return strings.TrimSpace(cleaned)
}

// ValidateEmail checks if an email string is structurally valid according to RFC 5322
func ValidateEmail(email string) bool {
	email = strings.TrimSpace(email)
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	addr, err := mail.ParseAddress(email)
	if err != nil || addr.Address != email {
		return false
	}
	// Check for double @ or spaces or header injection chars
	if strings.ContainsAny(email, "\r\n\t") {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return false
	}
	// Domain part must contain at least one dot
	if !strings.Contains(parts[1], ".") {
		return false
	}
	return true
}

// ValidateUsername checks username format
func ValidateUsername(username string) bool {
	return usernameRegex.MatchString(username)
}

// ValidateLength checks if text is between min and max UTF-8 characters
func ValidateLength(text string, min, max int) bool {
	length := utf8.RuneCountInString(text)
	return length >= min && length <= max
}
