package cache

import (
	"prometheus-test-task/internals/api"
	"sync"
)

type InMemoryCache struct {
	mu   *sync.Mutex
	data map[string]struct{} // map with empty struct for less memory usage, because empty struct don't use memory
}

func New() *InMemoryCache {
	return &InMemoryCache{
		mu:   &sync.Mutex{},
		data: make(map[string]struct{}),
	}
}

func (imc *InMemoryCache) Set(val string) error {
	imc.mu.Lock()
	defer imc.mu.Unlock()

	_, exists := imc.data[val]
	if exists {
		return api.ErrValExists{ // if exists value we returnin custom error
			Msg: val,
		}
	}

	imc.data[val] = struct{}{}
	return nil
}
