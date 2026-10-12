package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds all backend runtime configuration loaded from environment variables
type Config struct {
	// Server
	Port                string
	Environment         string // "development" or "production"
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	IdleTimeout         time.Duration
	ReadHeaderTimeout   time.Duration
	MaxHeaderBytes      int
	MaxRequestBodyBytes int64

	// Database
	DatabaseURL  string
	MaxOpenConns int
	MaxIdleConns int

	// CORS & Origins
	FrontendURL    string
	AllowedOrigins []string

	// Security & Auth
	CookieSecure       bool
	CookieDomain       string
	CookieSameSite     string
	SessionTTL         time.Duration
	AdminSessionSecret string
	CSRFSecret         string

	// Rate Limiting
	PublicRPM           int
	LoginRPM            int
	ContactRPM          int
	AdminRPM            int
	RateLimitGeneralReq int
	RateLimitContactReq int
	TrustProxy          bool

	// Network & Proxies
	TrustedProxies []*net.IPNet

	// Monitoring & Metrics
	MetricsEnabled bool

	// Email / Notifications (optional)
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	AdminEmail   string
}

// IsProduction returns true if the environment is production
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

// LoadConfig loads configuration from environment variables with sensible production defaults
func LoadConfig() (*Config, error) {
	env := getEnv("APP_ENV", "")
	if env == "" {
		env = getEnv("ENVIRONMENT", "development")
	}

	dbURL := getEnv("DATABASE_PATH", "")
	if dbURL == "" {
		dbURL = getEnv("DATABASE_URL", "./portfolio.db")
	}

	generalRPM := getEnvInt("RATE_LIMIT_GENERAL_REQ", 0)
	if generalRPM == 0 {
		generalRPM = getEnvInt("RATE_LIMIT_PUBLIC_RPM", 100)
	}

	contactRPM := getEnvInt("RATE_LIMIT_CONTACT_REQ", 0)
	if contactRPM == 0 {
		contactRPM = getEnvInt("RATE_LIMIT_CONTACT_RPM", 5)
	}

	trustProxy := getEnvBool("TRUST_PROXY", false)

	cfg := &Config{
		Port:                getEnv("PORT", "8080"),
		Environment:         env,
		ReadTimeout:         getEnvDuration("READ_TIMEOUT", 10*time.Second),
		WriteTimeout:        getEnvDuration("WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:         getEnvDuration("IDLE_TIMEOUT", 60*time.Second),
		ReadHeaderTimeout:   getEnvDuration("READ_HEADER_TIMEOUT", 5*time.Second),
		MaxHeaderBytes:      getEnvInt("MAX_HEADER_BYTES", 1048576),              // 1MB
		MaxRequestBodyBytes: int64(getEnvInt("MAX_REQUEST_BODY_BYTES", 1048576)), // 1MB

		DatabaseURL:  dbURL,
		MaxOpenConns: getEnvInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns: getEnvInt("DB_MAX_IDLE_CONNS", 10),

		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:5173"),

		CookieDomain:   getEnv("COOKIE_DOMAIN", ""),
		CookieSameSite: getEnv("COOKIE_SAME_SITE", "Lax"),
		SessionTTL:     getEnvDuration("ADMIN_SESSION_TTL", 24*time.Hour),

		PublicRPM:           generalRPM,
		LoginRPM:            getEnvInt("RATE_LIMIT_LOGIN_RPM", 5),
		ContactRPM:          contactRPM,
		AdminRPM:            getEnvInt("RATE_LIMIT_ADMIN_RPM", 60),
		RateLimitGeneralReq: generalRPM,
		RateLimitContactReq: contactRPM,
		TrustProxy:          trustProxy,

		MetricsEnabled: getEnvBool("METRICS_ENABLED", true),

		SMTPHost:     getEnv("SMTP_HOST", ""),
		SMTPPort:     getEnvInt("SMTP_PORT", 587),
		SMTPUsername: getEnv("SMTP_USERNAME", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
		AdminEmail:   getEnv("ADMIN_EMAIL", "valentineawili@gmail.com"),
	}

	// Cookie Secure default: true in production, false in development unless overridden
	cfg.CookieSecure = getEnvBool("COOKIE_SECURE", cfg.IsProduction())

	// Secrets
	cfg.AdminSessionSecret = getEnv("ADMIN_SESSION_SECRET", "dev-secret-session-key-change-in-prod-32bytes")
	cfg.CSRFSecret = getEnv("CSRF_SECRET", "dev-csrf-secret-key-change-in-prod-32bytes")

	// Allowed CORS origins
	originsRaw := getEnv("ALLOWED_ORIGINS", "")
	var origins []string
	if originsRaw != "" {
		for _, o := range strings.Split(originsRaw, ",") {
			trimmed := strings.TrimSpace(o)
			if trimmed != "" {
				origins = append(origins, trimmed)
			}
		}
	}
	if len(origins) == 0 {
		origins = []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
			"http://localhost:3000",
		}
		if cfg.FrontendURL != "" && !sliceContains(origins, cfg.FrontendURL) {
			origins = append(origins, cfg.FrontendURL)
		}
	}
	cfg.AllowedOrigins = origins

	// Parse trusted proxy CIDRs (e.g. "127.0.0.1/32,10.0.0.0/8,172.16.0.0/12,192.168.0.0/16")
	trustedRaw := getEnv("TRUSTED_PROXIES", "127.0.0.1/32,::1/128,10.0.0.0/8,172.16.0.0/12,192.168.0.0/16")
	for _, cidr := range strings.Split(trustedRaw, ",") {
		cidr = strings.TrimSpace(cidr)
		if cidr == "" {
			continue
		}
		if !strings.Contains(cidr, "/") {
			if strings.Contains(cidr, ":") {
				cidr = cidr + "/128"
			} else {
				cidr = cidr + "/32"
			}
		}
		_, ipNet, err := net.ParseCIDR(cidr)
		if err == nil {
			cfg.TrustedProxies = append(cfg.TrustedProxies, ipNet)
		}
	}

	return cfg, nil
}

// IsTrustedProxy checks whether a given client IP is an authorized reverse proxy (e.g. Nginx, Cloudflare)
func (c *Config) IsTrustedProxy(ipStr string) bool {
	if c.TrustProxy {
		return true
	}
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	for _, ipNet := range c.TrustedProxies {
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return strings.TrimSpace(val)
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		if parsed, err := strconv.Atoi(strings.TrimSpace(val)); err == nil {
			return parsed
		}
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		lower := strings.ToLower(strings.TrimSpace(val))
		return lower == "true" || lower == "1" || lower == "yes"
	}
	return defaultVal
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		if d, err := time.ParseDuration(strings.TrimSpace(val)); err == nil {
			return d
		}
		// If supplied as plain integer seconds
		if secs, err := strconv.Atoi(strings.TrimSpace(val)); err == nil {
			return time.Duration(secs) * time.Second
		}
	}
	return defaultVal
}

func sliceContains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

func (c *Config) Validate() error {
	if c.IsProduction() {
		if c.AdminSessionSecret == "dev-secret-session-key-change-in-prod-32bytes" {
			return fmt.Errorf("in production, ADMIN_SESSION_SECRET must be explicitly configured with a secure key")
		}
		if c.CSRFSecret == "dev-csrf-secret-key-change-in-prod-32bytes" {
			return fmt.Errorf("in production, CSRF_SECRET must be explicitly configured with a secure key")
		}
	}
	return nil
}
