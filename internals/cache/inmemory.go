package cache

import (
	"prometheus-test-task/internals/api"
	"sync"
)

type InMemoryCache struct {
	mu   *sync.Mutex
	data map[string]struct{}
}

func New() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]struct{}),
	}
}

func (imc *InMemoryCache) Set(val string) error {
	imc.mu.Lock()
	defer imc.mu.Unlock()

	_, exists := imc.data[val]
	if exists {
		return api.ErrValExists{
			Msg: val,
		}
	}

	imc.data[val] = struct{}{}
	return nil
}
