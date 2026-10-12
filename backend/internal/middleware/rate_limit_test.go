package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"portfolio-backend/internal/config"
)

func TestRateLimiter_TokenBucket(t *testing.T) {
	cfg := &config.Config{
		PublicRPM: 120, // 2 tokens per sec
	}

	// Create limiter with burst of 3
	limiter := NewIPRateLimiter(60, 3, cfg)
	defer limiter.Close()

	ip := "192.0.2.1"

	// 1. First 3 requests should be allowed immediately
	for i := 0; i < 3; i++ {
		if !limiter.Allow(ip) {
			t.Fatalf("expected request %d to be allowed", i+1)
		}
	}

	// 2. 4th immediate request should be rejected (burst exhausted)
	if limiter.Allow(ip) {
		t.Fatalf("expected 4th immediate request to be rejected")
	}

	// 3. Test HTTP middleware returns 429
	handler := limiter.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
	req.RemoteAddr = "192.0.2.1:12345"
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTooManyRequests {
		t.Fatalf("expected 429 Too Many Requests, got %d", rr.Code)
	}

	if rr.Header().Get("Retry-After") != "60" {
		t.Errorf("expected Retry-After 60, got %s", rr.Header().Get("Retry-After"))
	}
}
