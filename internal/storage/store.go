package store

import (
	"sync"
)

type MemStorage struct {
	sync.RWMutex
	items map[string]item
}

type item struct {
	Value interface{}
}

func (ms *MemStorage) Set(key string, value interface{}) {
	ms.Lock()
	defer ms.Unlock()

	ms.items[key] = item{
		Value: value,
	}
}

func (ms *MemStorage) Get(key string) (interface{}, bool) {
	ms.RLock()
	defer ms.RUnlock()

	item, found := ms.items[key]
	if !found {
		return nil, false
	}
	return item.Value, true
}

func (ms *MemStorage) Init() {
	ms.items = make(map[string]item)
}
