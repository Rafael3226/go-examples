package cache

import (
	"sync"
)

type SafeCache struct {
    mu sync.Mutex
    v map[string]string
}

func (sc *SafeCache) Set(key, value string) {
    sc.mu.
}