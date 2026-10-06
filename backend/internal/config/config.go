package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds all backend runtime configuration loaded from environment variables.
type Config struct {
	Port                   string
	AppEnv                 string
	DatabasePath           string
	AllowedOrigins         []string
	TrustProxy             bool
	RateLimitGeneralReq    int
	RateLimitGeneralWindow time.Duration
	RateLimitContactReq    int
	RateLimitContactWindow time.Duration
	ReadTimeout            time.Duration
	ReadHeaderTimeout      time.Duration
	WriteTimeout           time.Duration
	IdleTimeout            time.Duration
	MaxHeaderBytes         int
	MaxRequestBodyBytes    int64
}

// LoadConfig reads configuration from environment variables with safe defaults.
func LoadConfig() (*Config, error) {
	port := getEnv("PORT", "8080")
	appEnv := getEnv("APP_ENV", "development")

	dbPath := getEnv("DATABASE_PATH", "")
	if dbPath == "" {
		dbPath = getEnv("DATABASE_URL", "./data/portfolio.db")
	}

	originsRaw := getEnv("ALLOWED_ORIGINS", "")
	var allowedOrigins []string
	if originsRaw != "" {
		for _, o := range strings.Split(originsRaw, ",") {
			trimmed := strings.TrimSpace(o)
			if trimmed != "" {
				allowedOrigins = append(allowedOrigins, trimmed)
			}
		}
	} else {
		allowedOrigins = []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
			"http://localhost:3000",
		}
		if fe := os.Getenv("FRONTEND_URL"); fe != "" {
			allowedOrigins = append(allowedOrigins, strings.TrimSpace(fe))
		}
	}

	trustProxy := getEnvBool("TRUST_PROXY", false)
	rateLimitGeneralReq := getEnvInt("RATE_LIMIT_GENERAL_REQ", 100)
	rateLimitGeneralWindow := getEnvDuration("RATE_LIMIT_GENERAL_WINDOW", 1*time.Minute)
	rateLimitContactReq := getEnvInt("RATE_LIMIT_CONTACT_REQ", 5)
	rateLimitContactWindow := getEnvDuration("RATE_LIMIT_CONTACT_WINDOW", 10*time.Minute)

	readTimeout := getEnvDuration("READ_TIMEOUT", 10*time.Second)
	readHeaderTimeout := getEnvDuration("READ_HEADER_TIMEOUT", 5*time.Second)
	writeTimeout := getEnvDuration("WRITE_TIMEOUT", 15*time.Second)
	idleTimeout := getEnvDuration("IDLE_TIMEOUT", 60*time.Second)

	maxHeaderBytes := getEnvInt("MAX_HEADER_BYTES", 1<<20)            // 1MB
	maxReqBodyBytes := getEnvInt64("MAX_REQUEST_BODY_BYTES", 64*1024) // 64KB

	cfg := &Config{
		Port:                   port,
		AppEnv:                 appEnv,
		DatabasePath:           dbPath,
		AllowedOrigins:         allowedOrigins,
		TrustProxy:             trustProxy,
		RateLimitGeneralReq:    rateLimitGeneralReq,
		RateLimitGeneralWindow: rateLimitGeneralWindow,
		RateLimitContactReq:    rateLimitContactReq,
		RateLimitContactWindow: rateLimitContactWindow,
		ReadTimeout:            readTimeout,
		ReadHeaderTimeout:      readHeaderTimeout,
		WriteTimeout:           writeTimeout,
		IdleTimeout:            idleTimeout,
		MaxHeaderBytes:         maxHeaderBytes,
		MaxRequestBodyBytes:    maxReqBodyBytes,
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

// Validate checks critical configuration values.
func (c *Config) Validate() error {
	if c.Port == "" {
		return fmt.Errorf("PORT cannot be empty")
	}
	if c.DatabasePath == "" {
		return fmt.Errorf("DATABASE_PATH cannot be empty")
	}
	if len(c.AllowedOrigins) == 0 {
		return fmt.Errorf("at least one allowed origin must be configured")
	}
	if c.RateLimitGeneralReq <= 0 {
		return fmt.Errorf("RATE_LIMIT_GENERAL_REQ must be greater than zero")
	}
	if c.RateLimitContactReq <= 0 {
		return fmt.Errorf("RATE_LIMIT_CONTACT_REQ must be greater than zero")
	}
	return nil
}

// IsProduction returns true if running in production mode.
func (c *Config) IsProduction() bool {
	return strings.ToLower(c.AppEnv) == "production"
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}
	return b
}

func getEnvInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return i
}

func getEnvInt64(key string, defaultVal int64) int64 {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultVal
	}
	return i
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	d, err := time.ParseDuration(val)
	if err != nil {
		return defaultVal
	}
	return d
}
