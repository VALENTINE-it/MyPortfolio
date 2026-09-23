package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"
)

// RateLimiter limits the number of requests per IP within a sliding duration window
type RateLimiter struct {
	sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}

	// Periodic garbage collection for old entries
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			rl.cleanup()
		}
	}()

	return rl
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.Lock()
	defer rl.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.window)

	var validTimes []time.Time
	for _, t := range rl.requests[ip] {
		if t.After(cutoff) {
			validTimes = append(validTimes, t)
		}
	}

	if len(validTimes) >= rl.limit {
		rl.requests[ip] = validTimes
		return false
	}

	validTimes = append(validTimes, now)
	rl.requests[ip] = validTimes
	return true
}

func (rl *RateLimiter) cleanup() {
	rl.Lock()
	defer rl.Unlock()

	cutoff := time.Now().Add(-rl.window)
	for ip, times := range rl.requests {
		var valid []time.Time
		for _, t := range times {
			if t.After(cutoff) {
				valid = append(valid, t)
			}
		}
		if len(valid) == 0 {
			delete(rl.requests, ip)
		} else {
			rl.requests[ip] = valid
		}
	}
}

type ContactHandler struct {
	service     *services.ContactService
	rateLimiter *RateLimiter
}

func NewContactHandler(service *services.ContactService) *ContactHandler {
	// Allow 5 submissions per minute per IP
	return &ContactHandler{
		service:     service,
		rateLimiter: NewRateLimiter(5, 1*time.Minute),
	}
}

// HandleContact handles POST /api/contact
func (h *ContactHandler) HandleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	// 1. Rate limiting check
	ip := getClientIP(r)
	if !h.rateLimiter.Allow(ip) {
		writeJSON(w, http.StatusTooManyRequests, map[string]interface{}{
			"success": false,
			"message": "Too many requests. Please wait a minute before submitting again.",
		})
		return
	}

	// 2. Content-Type check
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(strings.ToLower(contentType), "application/json") {
		writeJSON(w, http.StatusUnsupportedMediaType, map[string]interface{}{
			"success": false,
			"message": "Content-Type must be application/json",
		})
		return
	}

	// 3. Limit request body size to 64KB
	r.Body = http.MaxBytesReader(w, r.Body, 64*1024)

	var req models.ContactRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Reject unexpected payload fields
	if err := decoder.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Malformed JSON payload or invalid request body format",
		})
		return
	}

	// 4. Validate and submit
	saved, err := h.service.SubmitContact(&req)
	if err != nil {
		var valErrs services.ValidationErrors
		if errors.As(err, &valErrs) {
			writeJSON(w, http.StatusBadRequest, map[string]interface{}{
				"success": false,
				"message": "Validation failed",
				"errors":  valErrs,
			})
			return
		}

		log.Printf("Internal error submitting contact: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": "Unable to process your request at this time. Please try again later.",
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Thank you for reaching out! Your message has been received.",
		"data": map[string]interface{}{
			"id":        saved.ID,
			"createdAt": saved.CreatedAt,
		},
	})
}

// getClientIP extracts real or proxy IP
func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		return strings.TrimSpace(xrip)
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
