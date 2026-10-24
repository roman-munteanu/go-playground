package ratelimit

import (
	"errors"
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	mu           sync.Mutex
	capacity     float64
	tokens       float64
	refillPerSec float64
	last         time.Time
}

func NewTokenBucket(capacity int, refillPerSecond float64) (*TokenBucket, error) {

	if capacity <= 0 || refillPerSecond <= 0 {
		return nil, errors.New("capacity and refill rate must be positive")
	}

	now := time.Now()

	return &TokenBucket{
		capacity:     float64(capacity),
		tokens:       float64(capacity),
		refillPerSec: refillPerSecond,
		last:         now,
	}, nil
}

func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.last).Seconds()

	if elapsed > 0 {
		b.tokens = math.Min(
			b.capacity,
			b.tokens+elapsed*b.refillPerSec,
		)
		b.last = now
	}

	if b.tokens < 1 {
		return false
	}

	b.tokens--
	return true
}
