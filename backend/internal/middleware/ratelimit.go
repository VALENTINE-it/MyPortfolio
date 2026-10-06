package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// RateLimiterStore defines an interface for rate limiting storage,
// allowing in-memory implementation to be swapped with Redis or external cache.
type RateLimiterStore interface {
	Allow(ctx context.Context, key string, limit int, window time.Duration) (allowed bool, retryAfter time.Duration, err error)
	Close() error
}

// MemoryRateLimiter implements RateLimiterStore using in-memory sliding window logs.
type MemoryRateLimiter struct {
	mu       sync.RWMutex
	requests map[string][]time.Time
	stopChan chan struct{}
}

// NewMemoryRateLimiter creates a new thread-safe memory rate limiter with background garbage collection.
func NewMemoryRateLimiter(cleanupInterval time.Duration) *MemoryRateLimiter {
	limiter := &MemoryRateLimiter{
		requests: make(map[string][]time.Time),
		stopChan: make(chan struct{}),
	}

	go limiter.cleanupLoop(cleanupInterval)

	return limiter
}

// Allow checks if the given key is allowed to perform an action within the limit and window.
func (m *MemoryRateLimiter) Allow(ctx context.Context, key string, limit int, window time.Duration) (bool, time.Duration, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-window)

	timestamps := m.requests[key]
	var valid []time.Time
	for _, t := range timestamps {
		if t.After(cutoff) {
			valid = append(valid, t)
		}
	}

	if len(valid) >= limit {
		m.requests[key] = valid
		oldest := valid[0]
		retryAfter := window - now.Sub(oldest)
		if retryAfter < 1*time.Second {
			retryAfter = 1 * time.Second
		}
		return false, retryAfter, nil
	}

	valid = append(valid, now)
	m.requests[key] = valid
	return true, 0, nil
}

// Close stops the background cleanup worker.
func (m *MemoryRateLimiter) Close() error {
	select {
	case <-m.stopChan:
		// already closed
	default:
		close(m.stopChan)
	}
	return nil
}

func (m *MemoryRateLimiter) cleanupLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-m.stopChan:
			return
		case <-ticker.C:
			m.cleanup(15 * time.Minute)
		}
	}
}

func (m *MemoryRateLimiter) cleanup(maxAge time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	cutoff := time.Now().Add(-maxAge)
	for key, timestamps := range m.requests {
		var valid []time.Time
		for _, t := range timestamps {
			if t.After(cutoff) {
				valid = append(valid, t)
			}
		}
		if len(valid) == 0 {
			delete(m.requests, key)
		} else {
			m.requests[key] = valid
		}
	}
}

// RateLimit creates an HTTP middleware that limits requests per IP based on the provided limit and window.
func RateLimit(store RateLimiterStore, limit int, window time.Duration, trustProxy bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP := GetClientIP(r, trustProxy)
			key := fmt.Sprintf("%s:%s", r.URL.Path, clientIP)

			allowed, retryAfter, err := store.Allow(r.Context(), key, limit, window)
			if err != nil {
				slog.Error("Rate limiter error", "error", err, "ip", clientIP)
				next.ServeHTTP(w, r)
				return
			}

			if !allowed {
				retrySeconds := int(retryAfter.Seconds())
				if retrySeconds <= 0 {
					retrySeconds = 1
				}

				reqID := GetRequestID(r.Context())
				slog.Warn("Rate limit exceeded",
					"request_id", reqID,
					"ip", clientIP,
					"path", r.URL.Path,
					"retry_after_sec", retrySeconds,
				)

				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Retry-After", fmt.Sprintf("%d", retrySeconds))
				w.WriteHeader(http.StatusTooManyRequests)
				_ = json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"message": fmt.Sprintf("Too many requests. Please retry in %d seconds.", retrySeconds),
					"error": map[string]interface{}{
						"code":    "RATE_LIMIT_EXCEEDED",
						"message": fmt.Sprintf("Too many requests. Please retry in %d seconds.", retrySeconds),
					},
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
