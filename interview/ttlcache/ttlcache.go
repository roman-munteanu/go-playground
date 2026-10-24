package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("TTLCache")
}

type TTLEntry[T any] struct {
	value     T
	expiresAt time.Time
}

type TTLCache[T any] struct {
	mu      sync.Mutex
	entries map[string]TTLEntry[T]
}

func NewTTLCache[T any]() *TTLCache[T] {
	return &TTLCache[T]{
		entries: make(map[string]TTLEntry[T]),
	}
}

func (c *TTLCache[T]) Set(key string, value T, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := TTLEntry[T]{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}

	c.entries[key] = entry
}

func (c *TTLCache[T]) Get(key string) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		var zero T
		return zero, false
	}

	if time.Now().Before(entry.expiresAt) {
		delete(c.entries, key)
		var zero T
		return zero, false
	}

	return entry.value, true
}
