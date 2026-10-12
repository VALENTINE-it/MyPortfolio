package cache

import (
	"context"
	"time"
)

// Cache defines the caching interface to allow swapping between Memory and Redis
type Cache interface {
	Get(ctx context.Context, key string) ([]byte, bool)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration)
	Delete(ctx context.Context, key string)
	DeletePrefix(ctx context.Context, prefix string)
	Flush(ctx context.Context)
}
