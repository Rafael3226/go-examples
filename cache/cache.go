package cache

import (
	"sync"
)

type SafeCache[T any] struct {
	mu    sync.Mutex
	value map[string]T
}

type IMap[T any] interface {
	Set(key string, value T)
	Get(key string) (T, bool)
}

func (sc *SafeCache[T]) Set(key string, value T) {
	sc.mu.Lock()
	sc.value[key] = value
	sc.mu.Unlock()
}

func (sc *SafeCache[T]) Get(key string) (T, bool) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	val, found := sc.value[key]
	return val, found
}

func CreateSafeCache[T any]() SafeCache[T] {
	return SafeCache[T]{value: make(map[string]T)}
}
