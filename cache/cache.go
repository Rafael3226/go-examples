package cache

import (
	"sync"
)

type SafeCache[T any] struct {
	mu    sync.Mutex
	value map[string]T
}

func (sc *SafeCache[T]) Set(key string, value T) {
	sc.mu.Lock()
	sc.value[key] = value
	sc.mu.Unlock()
}

func (sc *SafeCache[T]) Get(key string) T {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.value[key]
}
