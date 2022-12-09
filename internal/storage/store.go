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

// Установка значений
func (c *MemStorage) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = item{
		Value: value,
	}
}

// Получение значения из хранилища по ключу
func (c *MemStorage) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}
	return item.Value, true
}

// Инициализация хранилища нужными метриками
func (c *MemStorage) Init() {
	c.items = make(map[string]item)
	c.Set("Alloc", metrics.Gauge(0))
	c.Set("BuckHashSys", metrics.Gauge(0))
	c.Set("Frees", metrics.Gauge(0))
	c.Set("GCCPUFraction", metrics.Gauge(0))
	c.Set("GCSys", metrics.Gauge(0))
	c.Set("HeapAlloc", metrics.Gauge(0))
	c.Set("HeapIdle", metrics.Gauge(0))
	c.Set("HeapInuse", metrics.Gauge(0))
	c.Set("HeapObjects", metrics.Gauge(0))
	c.Set("HeapReleased", metrics.Gauge(0))
	c.Set("HeapSys", metrics.Gauge(0))
	c.Set("LastGC", metrics.Gauge(0))
	c.Set("Lookups", metrics.Gauge(0))
	c.Set("MCacheInuse", metrics.Gauge(0))
	c.Set("MCacheSys", metrics.Gauge(0))
	c.Set("MSpanInuse", metrics.Gauge(0))
	c.Set("MSpanSys", metrics.Gauge(0))
	c.Set("Mallocs", metrics.Gauge(0))
	c.Set("NextGC", metrics.Gauge(0))
	c.Set("NumForcedGC", metrics.Gauge(0))
	c.Set("NumGC", metrics.Gauge(0))
	c.Set("OtherSys", metrics.Gauge(0))
	c.Set("PauseTotalNs", metrics.Gauge(0))
	c.Set("StackInuse", metrics.Gauge(0))
	c.Set("StackSys", metrics.Gauge(0))
	c.Set("Sys", metrics.Gauge(0))
	c.Set("TotalAlloc", metrics.Gauge(0))
	c.Set("RandomValue", metrics.Gauge(0))
	c.Set("PollCount", metrics.Counter(0))
}
