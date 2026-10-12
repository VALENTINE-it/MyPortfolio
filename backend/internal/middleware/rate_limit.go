package middleware

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"portfolio-backend/internal/config"
	"portfolio-backend/internal/monitoring"
	"portfolio-backend/internal/utils"
)

// tokenBucket tracks available tokens for a specific client IP
type tokenBucket struct {
	tokens     float64
	lastRefill time.Time
}

// IPRateLimiter implements a concurrency-safe, memory-bounded token bucket rate limiter
type IPRateLimiter struct {
	mu         sync.Mutex
	buckets    map[string]*tokenBucket
	rate       float64       // tokens added per second
	burst      int           // maximum bucket capacity
	window     time.Duration // idle cleanup threshold
	stopGC     chan struct{}
	cfg        *config.Config
}

// NewIPRateLimiter creates a new rate limiter with specified requests-per-minute limit
func NewIPRateLimiter(rpm int, burst int, cfg *config.Config) *IPRateLimiter {
	if rpm <= 0 {
		rpm = 60
	}
	if burst <= 0 {
		burst = rpm / 2
		if burst < 5 {
			burst = 5
		}
	}

	rate := float64(rpm) / 60.0 // tokens per second

	rl := &IPRateLimiter{
		buckets: make(map[string]*tokenBucket),
		rate:    rate,
		burst:   burst,
		window:  10 * time.Minute,
		stopGC:  make(chan struct{}),
		cfg:     cfg,
	}

	go rl.startCleanup(2 * time.Minute)

	return rl
}

func (rl *IPRateLimiter) Close() {
	close(rl.stopGC)
}

// Allow checks if the given IP has tokens available
func (rl *IPRateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	b, exists := rl.buckets[ip]
	if !exists {
		// Enforce memory bounds: cap at 10,000 active tracked IPs
		if len(rl.buckets) >= 10000 {
			rl.cleanupExpired(now)
		}

		rl.buckets[ip] = &tokenBucket{
			tokens:     float64(rl.burst - 1),
			lastRefill: now,
		}
		return true
	}

	// Refill tokens based on elapsed time
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.tokens += elapsed * rl.rate
	if b.tokens > float64(rl.burst) {
		b.tokens = float64(rl.burst)
	}
	b.lastRefill = now

	if b.tokens >= 1.0 {
		b.tokens -= 1.0
		return true
	}

	return false
}

// GetClientIP securely extracts client IP taking trusted reverse proxies into account
func (rl *IPRateLimiter) GetClientIP(r *http.Request) string {
	remoteHost, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		remoteHost = r.RemoteAddr
	}

	// Only trust forwarded headers if the immediate peer is a configured trusted reverse proxy
	isTrusted := false
	if rl.cfg != nil {
		isTrusted = rl.cfg.IsTrustedProxy(remoteHost)
	}

	if isTrusted {
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			parts := strings.Split(xff, ",")
			clientIP := strings.TrimSpace(parts[0])
			if net.ParseIP(clientIP) != nil {
				return clientIP
			}
		}
		if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
			clientIP := strings.TrimSpace(xrip)
			if net.ParseIP(clientIP) != nil {
				return clientIP
			}
		}
	}

	return remoteHost
}

// Middleware creates an HTTP middleware handler enforcing rate limits
func (rl *IPRateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := rl.GetClientIP(r)

		if !rl.Allow(ip) {
			monitoring.GetMetrics().RecordRateLimitRejected()
			w.Header().Set("Retry-After", "60")
			utils.SendError(w, http.StatusTooManyRequests, utils.CodeTooManyRequests, fmt.Sprintf("Too many requests from IP %s. Please retry in 60 seconds.", ip))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rl *IPRateLimiter) startCleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-rl.stopGC:
			return
		case <-ticker.C:
			rl.mu.Lock()
			rl.cleanupExpired(time.Now())
			rl.mu.Unlock()
		}
	}
}

func (rl *IPRateLimiter) cleanupExpired(now time.Time) {
	cutoff := now.Add(-rl.window)
	for ip, b := range rl.buckets {
		if b.lastRefill.Before(cutoff) {
			delete(rl.buckets, ip)
		}
	}
}
