package middleware

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestIDMiddleware(t *testing.T) {
	// 1. Without incoming request ID -> generates new one
	handler := RequestIDMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := GetRequestID(r.Context())
		if reqID == "" {
			t.Errorf("expected request ID in context, got empty")
		}
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	headerID := rr.Header().Get("X-Request-ID")
	if headerID == "" {
		t.Errorf("expected X-Request-ID header in response")
	}

	// 2. With incoming request ID -> preserves it
	reqWithID := httptest.NewRequest(http.MethodGet, "/test", nil)
	reqWithID.Header.Set("X-Request-ID", "custom-trace-id-1234")
	rrWithID := httptest.NewRecorder()
	handler.ServeHTTP(rrWithID, reqWithID)

	if got := rrWithID.Header().Get("X-Request-ID"); got != "custom-trace-id-1234" {
		t.Errorf("expected X-Request-ID 'custom-trace-id-1234', got '%s'", got)
	}
}

func TestSecurityHeadersMiddleware(t *testing.T) {
	// 1. Non-production (no HSTS)
	handlerDev := SecurityHeaders(false, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rrDev := httptest.NewRecorder()
	handlerDev.ServeHTTP(rrDev, req)

	if rrDev.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Errorf("expected X-Content-Type-Options: nosniff")
	}
	if rrDev.Header().Get("X-Frame-Options") != "DENY" {
		t.Errorf("expected X-Frame-Options: DENY")
	}
	if rrDev.Header().Get("Referrer-Policy") != "strict-origin-when-cross-origin" {
		t.Errorf("expected Referrer-Policy: strict-origin-when-cross-origin")
	}
	if rrDev.Header().Get("Strict-Transport-Security") != "" {
		t.Errorf("HSTS should not be set in dev mode")
	}

	// 2. Production (with HSTS)
	handlerProd := SecurityHeaders(true, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	rrProd := httptest.NewRecorder()
	handlerProd.ServeHTTP(rrProd, req)

	if rrProd.Header().Get("Strict-Transport-Security") == "" {
		t.Errorf("expected Strict-Transport-Security in prod mode")
	}
}

func TestCORSMiddleware(t *testing.T) {
	allowedOrigins := []string{"https://example.com", "http://localhost:5173"}
	handler := CORS(allowedOrigins, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// 1. Allowed origin GET
	req := httptest.NewRequest(http.MethodGet, "/api/projects", nil)
	req.Header.Set("Origin", "https://example.com")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
	if rr.Header().Get("Access-Control-Allow-Origin") != "https://example.com" {
		t.Errorf("expected Access-Control-Allow-Origin: https://example.com, got %s", rr.Header().Get("Access-Control-Allow-Origin"))
	}

	// 2. Preflight OPTIONS for allowed origin
	reqOpt := httptest.NewRequest(http.MethodOptions, "/api/projects", nil)
	reqOpt.Header.Set("Origin", "http://localhost:5173")
	rrOpt := httptest.NewRecorder()
	handler.ServeHTTP(rrOpt, reqOpt)

	if rrOpt.Code != http.StatusNoContent {
		t.Errorf("expected 204 No Content for preflight, got %d", rrOpt.Code)
	}

	// 3. Preflight OPTIONS for disallowed origin
	reqBadOpt := httptest.NewRequest(http.MethodOptions, "/api/projects", nil)
	reqBadOpt.Header.Set("Origin", "https://malicious.com")
	rrBadOpt := httptest.NewRecorder()
	handler.ServeHTTP(rrBadOpt, reqBadOpt)

	if rrBadOpt.Code != http.StatusForbidden {
		t.Errorf("expected 403 Forbidden for unauthorized preflight, got %d", rrBadOpt.Code)
	}

	// 4. Request without Origin (e.g. server-to-server or curl)
	reqNoOrigin := httptest.NewRequest(http.MethodGet, "/api/projects", nil)
	rrNoOrigin := httptest.NewRecorder()
	handler.ServeHTTP(rrNoOrigin, reqNoOrigin)

	if rrNoOrigin.Code != http.StatusOK {
		t.Errorf("expected 200 for request without origin header, got %d", rrNoOrigin.Code)
	}
}

func TestPanicRecoveryMiddleware(t *testing.T) {
	panickingHandler := PanicRecovery(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("unexpected runtime panic")
	}))

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rr := httptest.NewRecorder()

	// Should not crash the test process
	panickingHandler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 status on panic, got %d", rr.Code)
	}
	if !bytes.Contains(rr.Body.Bytes(), []byte("An internal server error occurred.")) {
		t.Errorf("expected generic error message in response body")
	}
}

func TestClientIPExtraction(t *testing.T) {
	// 1. Untrusted proxy -> strictly uses RemoteAddr, ignores X-Forwarded-For
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.RemoteAddr = "192.168.1.50:1234"
	req.Header.Set("X-Forwarded-For", "8.8.8.8, 10.0.0.1")

	ipUntrusted := GetClientIP(req, false)
	if ipUntrusted != "192.168.1.50" {
		t.Errorf("expected remote addr IP '192.168.1.50', got '%s'", ipUntrusted)
	}

	// 2. Trusted proxy -> extracts first client IP
	ipTrusted := GetClientIP(req, true)
	if ipTrusted != "8.8.8.8" {
		t.Errorf("expected client IP '8.8.8.8' when trusting proxy, got '%s'", ipTrusted)
	}
}

func TestRateLimiterMiddleware(t *testing.T) {
	limiter := NewMemoryRateLimiter(1 * time.Minute)
	defer limiter.Close()

	middlewareFunc := RateLimit(limiter, 2, 500*time.Millisecond, false)
	handler := middlewareFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.RemoteAddr = "10.0.0.5:1234"

	// Request 1: OK
	rr1 := httptest.NewRecorder()
	handler.ServeHTTP(rr1, req)
	if rr1.Code != http.StatusOK {
		t.Fatalf("request 1 should succeed, got %d", rr1.Code)
	}

	// Request 2: OK
	rr2 := httptest.NewRecorder()
	handler.ServeHTTP(rr2, req)
	if rr2.Code != http.StatusOK {
		t.Fatalf("request 2 should succeed, got %d", rr2.Code)
	}

	// Request 3: Rejected (429)
	rr3 := httptest.NewRecorder()
	handler.ServeHTTP(rr3, req)
	if rr3.Code != http.StatusTooManyRequests {
		t.Fatalf("request 3 should be rate limited (429), got %d", rr3.Code)
	}
	if rr3.Header().Get("Retry-After") == "" {
		t.Errorf("expected Retry-After header on 429 response")
	}

	// Wait for window to expire
	time.Sleep(600 * time.Millisecond)

	// Request 4: OK again
	rr4 := httptest.NewRecorder()
	handler.ServeHTTP(rr4, req)
	if rr4.Code != http.StatusOK {
		t.Fatalf("request 4 should succeed after window reset, got %d", rr4.Code)
	}
}

func TestBodyLimitMiddleware(t *testing.T) {
	maxBytes := int64(10)
	middlewareFunc := BodyLimit(maxBytes)
	handler := middlewareFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "body too large", http.StatusRequestEntityTooLarge)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))

	// 1. Small body -> OK
	smallReq := httptest.NewRequest(http.MethodPost, "/upload", bytes.NewBufferString("short"))
	smallRr := httptest.NewRecorder()
	handler.ServeHTTP(smallRr, smallReq)
	if smallRr.Code != http.StatusOK {
		t.Errorf("expected small body to succeed, got %d", smallRr.Code)
	}

	// 2. Large body -> MaxBytesReader triggers error
	largeReq := httptest.NewRequest(http.MethodPost, "/upload", bytes.NewBufferString("this payload exceeds 10 bytes"))
	largeRr := httptest.NewRecorder()
	handler.ServeHTTP(largeRr, largeReq)
	if largeRr.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("expected large body to return 413, got %d", largeRr.Code)
	}
}
