package common

import (
	"sync"
	"sync/atomic"
	"time"
)

type rateLimitEntry struct {
	mu    sync.Mutex
	queue []int64
}

type InMemoryRateLimiter struct {
	store              sync.Map // map[string]*rateLimitEntry
	expirationDuration time.Duration
	initialized        int32 // atomic flag
}

func (l *InMemoryRateLimiter) Init(expirationDuration time.Duration) {
	if atomic.CompareAndSwapInt32(&l.initialized, 0, 1) {
		l.expirationDuration = expirationDuration
		if expirationDuration > 0 {
			go l.clearExpiredItems()
		}
	}
}

func (l *InMemoryRateLimiter) clearExpiredItems() {
	for {
		time.Sleep(l.expirationDuration)
		now := time.Now().Unix()
		l.store.Range(func(key, value any) bool {
			entry := value.(*rateLimitEntry)
			entry.mu.Lock()
			size := len(entry.queue)
			if size == 0 || now-entry.queue[size-1] > int64(l.expirationDuration.Seconds()) {
				entry.mu.Unlock()
				l.store.Delete(key)
			} else {
				entry.mu.Unlock()
			}
			return true
		})
	}
}

// Request parameter duration's unit is seconds
func (l *InMemoryRateLimiter) Request(key string, maxRequestNum int, duration int64) bool {
	val, _ := l.store.LoadOrStore(key, &rateLimitEntry{
		queue: make([]int64, 0, maxRequestNum),
	})
	entry := val.(*rateLimitEntry)
	entry.mu.Lock()
	defer entry.mu.Unlock()

	now := time.Now().Unix()
	if len(entry.queue) < maxRequestNum {
		entry.queue = append(entry.queue, now)
		return true
	}
	if now-entry.queue[0] >= duration {
		entry.queue = entry.queue[1:]
		entry.queue = append(entry.queue, now)
		return true
	}
	return false
}
