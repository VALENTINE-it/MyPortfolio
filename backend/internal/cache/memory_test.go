package cache

import (
	"context"
	"testing"
	"time"
)

func TestMemoryCache_SetGetExpire(t *testing.T) {
	c := NewMemoryCache(10, 50*time.Millisecond)
	defer c.Close()

	ctx := context.Background()

	// 1. Set and Get
	c.Set(ctx, "key1", []byte("value1"), 50*time.Millisecond)
	val, ok := c.Get(ctx, "key1")
	if !ok || string(val) != "value1" {
		t.Fatalf("expected to get value1, got %s (ok=%v)", string(val), ok)
	}

	// 2. Wait for expiration
	time.Sleep(70 * time.Millisecond)
	_, ok = c.Get(ctx, "key1")
	if ok {
		t.Fatalf("expected key1 to expire, but it was still found")
	}
}

func TestMemoryCache_DeleteAndPrefix(t *testing.T) {
	c := NewMemoryCache(10, 5*time.Minute)
	defer c.Close()

	ctx := context.Background()
	c.Set(ctx, "skills:1", []byte("s1"), 0)
	c.Set(ctx, "skills:2", []byte("s2"), 0)
	c.Set(ctx, "projects:1", []byte("p1"), 0)

	// Delete prefix
	c.DeletePrefix(ctx, "skills:")

	if _, ok := c.Get(ctx, "skills:1"); ok {
		t.Fatalf("expected skills:1 to be deleted")
	}
	if _, ok := c.Get(ctx, "skills:2"); ok {
		t.Fatalf("expected skills:2 to be deleted")
	}
	if _, ok := c.Get(ctx, "projects:1"); !ok {
		t.Fatalf("expected projects:1 to remain in cache")
	}
}
