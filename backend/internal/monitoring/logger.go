package monitoring

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type contextKey string

const (
	RequestIDKey contextKey = "request_id"
)

// InitLogger sets up log/slog with JSON handler in production or Text handler in development
func InitLogger(env string, out io.Writer) *slog.Logger {
	if out == nil {
		out = os.Stdout
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if env == "production" {
		handler = slog.NewJSONHandler(out, opts)
	} else {
		handler = slog.NewTextHandler(out, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}

// GetRequestID extracts Request ID from context if available
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if id, ok := ctx.Value(RequestIDKey).(string); ok {
		return id
	}
	return ""
}

// WithRequestID returns a new context carrying the given request ID
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

// LogWithContext logs message with request ID attribute attached if present
func LogWithContext(ctx context.Context, logger *slog.Logger) *slog.Logger {
	if logger == nil {
		logger = slog.Default()
	}
	reqID := GetRequestID(ctx)
	if reqID != "" {
		return logger.With("request_id", reqID)
	}
	return logger
}
