package cache

import (
	"context"
	"strings"
	"sync"
	"time"

	"portfolio-backend/internal/monitoring"
)

type cacheItem struct {
	value     []byte
	expiresAt time.Time
}

// MemoryCache is a concurrency-safe in-memory cache with TTL expiration and bounded entries
type MemoryCache struct {
	mu         sync.RWMutex
	items      map[string]cacheItem
	maxEntries int
	defaultTTL time.Duration
	stopGC     chan struct{}
}

// NewMemoryCache creates a new in-memory cache with background cleanup
func NewMemoryCache(maxEntries int, defaultTTL time.Duration) *MemoryCache {
	if maxEntries <= 0 {
		maxEntries = 1000
	}
	if defaultTTL <= 0 {
		defaultTTL = 5 * time.Minute
	}

	mc := &MemoryCache{
		items:      make(map[string]cacheItem, maxEntries),
		maxEntries: maxEntries,
		defaultTTL: defaultTTL,
		stopGC:     make(chan struct{}),
	}

	go mc.startPruning(1 * time.Minute)

	return mc
}

func (mc *MemoryCache) Close() {
	close(mc.stopGC)
}

func (mc *MemoryCache) Get(ctx context.Context, key string) ([]byte, bool) {
	mc.mu.RLock()
	item, found := mc.items[key]
	mc.mu.RUnlock()

	if !found {
		monitoring.GetMetrics().RecordCacheMiss()
		return nil, false
	}

	if time.Now().After(item.expiresAt) {
		mc.mu.Lock()
		delete(mc.items, key)
		mc.mu.Unlock()
		monitoring.GetMetrics().RecordCacheMiss()
		return nil, false
	}

	monitoring.GetMetrics().RecordCacheHit()
	return item.value, true
}

func (mc *MemoryCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) {
	if ttl <= 0 {
		ttl = mc.defaultTTL
	}

	mc.mu.Lock()
	defer mc.mu.Unlock()

	// If at capacity, evict an arbitrary expired entry or oldest entry
	if len(mc.items) >= mc.maxEntries {
		now := time.Now()
		var evictedKey string
		for k, v := range mc.items {
			if now.After(v.expiresAt) {
				evictedKey = k
				break
			}
			evictedKey = k
		}
		if evictedKey != "" {
			delete(mc.items, evictedKey)
		}
	}

	mc.items[key] = cacheItem{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
}

func (mc *MemoryCache) Delete(ctx context.Context, key string) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	delete(mc.items, key)
}

func (mc *MemoryCache) DeletePrefix(ctx context.Context, prefix string) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	for k := range mc.items {
		if strings.HasPrefix(k, prefix) {
			delete(mc.items, k)
		}
	}
}

func (mc *MemoryCache) Flush(ctx context.Context) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.items = make(map[string]cacheItem, mc.maxEntries)
}

func (mc *MemoryCache) startPruning(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-mc.stopGC:
			return
		case <-ticker.C:
			mc.pruneExpired()
		}
	}
}

func (mc *MemoryCache) pruneExpired() {
	now := time.Now()
	mc.mu.Lock()
	defer mc.mu.Unlock()

	for k, v := range mc.items {
		if now.After(v.expiresAt) {
			delete(mc.items, k)
		}
	}
}
