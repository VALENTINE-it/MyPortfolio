package config

import (
	"os"
	"testing"
	"time"
)

func TestConfig_Defaults(t *testing.T) {
	// Clear any overrides
	os.Unsetenv("PORT")
	os.Unsetenv("DATABASE_PATH")
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("ALLOWED_ORIGINS")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("expected valid config, got error: %v", err)
	}

	if cfg.Port != "8080" {
		t.Errorf("expected default port 8080, got %s", cfg.Port)
	}
	if cfg.RateLimitGeneralReq != 100 {
		t.Errorf("expected default general rate limit 100, got %d", cfg.RateLimitGeneralReq)
	}
	if cfg.RateLimitContactReq != 5 {
		t.Errorf("expected default contact rate limit 5, got %d", cfg.RateLimitContactReq)
	}
	if len(cfg.AllowedOrigins) == 0 {
		t.Errorf("expected default allowed origins to be non-empty")
	}
	if cfg.ReadTimeout != 10*time.Second {
		t.Errorf("expected default read timeout 10s, got %v", cfg.ReadTimeout)
	}
}

func TestConfig_CustomEnvironment(t *testing.T) {
	os.Setenv("PORT", "9090")
	os.Setenv("APP_ENV", "production")
	os.Setenv("DATABASE_PATH", "./custom.db")
	os.Setenv("ALLOWED_ORIGINS", "https://example.com,https://portfolio.com")
	os.Setenv("TRUST_PROXY", "true")
	os.Setenv("RATE_LIMIT_GENERAL_REQ", "50")
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("APP_ENV")
		os.Unsetenv("DATABASE_PATH")
		os.Unsetenv("ALLOWED_ORIGINS")
		os.Unsetenv("TRUST_PROXY")
		os.Unsetenv("RATE_LIMIT_GENERAL_REQ")
	}()

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Port != "9090" {
		t.Errorf("expected port 9090, got %s", cfg.Port)
	}
	if !cfg.IsProduction() {
		t.Errorf("expected IsProduction() to be true")
	}
	if !cfg.TrustProxy {
		t.Errorf("expected TrustProxy to be true")
	}
	if cfg.RateLimitGeneralReq != 50 {
		t.Errorf("expected general rate limit 50, got %d", cfg.RateLimitGeneralReq)
	}
	if len(cfg.AllowedOrigins) != 2 {
		t.Errorf("expected 2 allowed origins, got %d", len(cfg.AllowedOrigins))
	}
}
