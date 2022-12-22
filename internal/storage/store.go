package store

import (
	"devops/internal/metrics"
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
	ms.Set("Alloc", metrics.Gauge(0))
	ms.Set("BuckHashSys", metrics.Gauge(0))
	ms.Set("Frees", metrics.Gauge(0))
	ms.Set("GCCPUFraction", metrics.Gauge(0))
	ms.Set("GCSys", metrics.Gauge(0))
	ms.Set("HeapAlloc", metrics.Gauge(0))
	ms.Set("HeapIdle", metrics.Gauge(0))
	ms.Set("HeapInuse", metrics.Gauge(0))
	ms.Set("HeapObjects", metrics.Gauge(0))
	ms.Set("HeapReleased", metrics.Gauge(0))
	ms.Set("HeapSys", metrics.Gauge(0))
	ms.Set("LastGC", metrics.Gauge(0))
	ms.Set("Lookups", metrics.Gauge(0))
	ms.Set("MCacheInuse", metrics.Gauge(0))
	ms.Set("MCacheSys", metrics.Gauge(0))
	ms.Set("MSpanInuse", metrics.Gauge(0))
	ms.Set("MSpanSys", metrics.Gauge(0))
	ms.Set("Mallocs", metrics.Gauge(0))
	ms.Set("NextGC", metrics.Gauge(0))
	ms.Set("NumForcedGC", metrics.Gauge(0))
	ms.Set("NumGC", metrics.Gauge(0))
	ms.Set("OtherSys", metrics.Gauge(0))
	ms.Set("PauseTotalNs", metrics.Gauge(0))
	ms.Set("StackInuse", metrics.Gauge(0))
	ms.Set("StackSys", metrics.Gauge(0))
	ms.Set("Sys", metrics.Gauge(0))
	ms.Set("TotalAlloc", metrics.Gauge(0))
	ms.Set("RandomValue", metrics.Gauge(0))
	ms.Set("PollCount", metrics.Counter(0))
}
