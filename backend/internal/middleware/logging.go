package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"portfolio-backend/internal/monitoring"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriterWrapper) Write(b []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}
	n, err := rw.ResponseWriter.Write(b)
	rw.bytesWritten += n
	return n, err
}

// Logging logs all HTTP requests structured with log/slog and feeds telemetry to metrics
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID := monitoring.GetRequestID(r.Context())

		metrics := monitoring.GetMetrics()
		metrics.IncActiveRequests()
		defer metrics.DecActiveRequests()

		wrapper := &responseWriterWrapper{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)

		duration := time.Since(start)

		// Record in metrics
		metrics.RecordRequest(r.Method, r.URL.Path, wrapper.statusCode, duration)

		// Don't clutter logs with high-frequency health probes in development unless they fail
		if r.URL.Path == "/health/live" && wrapper.statusCode == http.StatusOK {
			return
		}

		logLevel := slog.LevelInfo
		if wrapper.statusCode >= 500 {
			logLevel = slog.LevelError
		} else if wrapper.statusCode >= 400 {
			logLevel = slog.LevelWarn
		}

		slog.Log(r.Context(), logLevel, "HTTP request handled",
			"method", r.Method,
			"path", r.URL.Path,
			"status", wrapper.statusCode,
			"duration_ms", float64(duration.Microseconds())/1000.0,
			"bytes", wrapper.bytesWritten,
			"request_id", reqID,
			"remote_ip", r.RemoteAddr,
		)
	})
}
