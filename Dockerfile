# ==========================================
# Multi-stage production Dockerfile
# ==========================================

# 1. Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /build

# Install CA certificates and tzdata
RUN apk add --no-cache ca-certificates tzdata

# Download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy backend source code
COPY backend/ ./

# Compile binary statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /build/bin/server ./cmd/server && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /build/bin/admin-cli ./cmd/admin-cli

# 2. Runtime stage
FROM alpine:3.20

# Create unprivileged application user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Create persistent database & backups directory
RUN mkdir -p /data /backups && chown -R appuser:appgroup /data /backups

WORKDIR /app

# Copy CA certs, tzdata, and compiled binaries from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder --chown=appuser:appgroup /build/bin/server /app/server
COPY --from=builder --chown=appuser:appgroup /build/bin/admin-cli /app/admin-cli

# Switch to non-root user
USER appuser

# Expose HTTP port
EXPOSE 8080

# Configure persistent SQLite path
ENV DATABASE_URL=/data/portfolio.db \
    PORT=8080 \
    ENVIRONMENT=production

# Health check
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD wget -qO- http://localhost:8080/health/live || exit 1

ENTRYPOINT ["/app/server"]
