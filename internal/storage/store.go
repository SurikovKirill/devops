package store

import (
	"devops/internal/helpers"
	"sync"
)

type MemStorage struct {
	sync.RWMutex
	items map[string]Item
}

type Item struct {
	Value interface{}
}

// Создаем новый экземпляр хранилища
func New() *MemStorage {
	// инициализируем карту(map) в паре ключ(string)/значение(Item)
	items := make(map[string]Item)

	cache := MemStorage{
		items: items,
	}
	return &cache
}

// Установка значений
func (c *MemStorage) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = Item{
		Value: value,
	}
}

// Получение значения из хранилища по ключу
func (c *MemStorage) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]
	// ключ не найден
	if !found {
		return nil, false
	}
	return item.Value, true
}

// Инициализация хранилища нужными метриками
func (c *MemStorage) Init() {
	c.Set("Alloc", helpers.Gauge(0))
	c.Set("BuckHashSys", helpers.Gauge(0))
	c.Set("Frees", helpers.Gauge(0))
	c.Set("GCCPUFraction", helpers.Gauge(0))
	c.Set("GCSys", helpers.Gauge(0))
	c.Set("HeapAlloc", helpers.Gauge(0))
	c.Set("HeapIdle", helpers.Gauge(0))
	c.Set("HeapInuse", helpers.Gauge(0))
	c.Set("HeapObjects", helpers.Gauge(0))
	c.Set("HeapReleased", helpers.Gauge(0))
	c.Set("HeapSys", helpers.Gauge(0))
	c.Set("LastGC", helpers.Gauge(0))
	c.Set("Lookups", helpers.Gauge(0))
	c.Set("MCacheInuse", helpers.Gauge(0))
	c.Set("MCacheSys", helpers.Gauge(0))
	c.Set("MSpanInuse", helpers.Gauge(0))
	c.Set("MSpanSys", helpers.Gauge(0))
	c.Set("Mallocs", helpers.Gauge(0))
	c.Set("NextGC", helpers.Gauge(0))
	c.Set("NumForcedGC", helpers.Gauge(0))
	c.Set("NumGC", helpers.Gauge(0))
	c.Set("OtherSys", helpers.Gauge(0))
	c.Set("PauseTotalNs", helpers.Gauge(0))
	c.Set("StackInuse", helpers.Gauge(0))
	c.Set("StackSys", helpers.Gauge(0))
	c.Set("Sys", helpers.Gauge(0))
	c.Set("TotalAlloc", helpers.Gauge(0))
	c.Set("RandomValue", helpers.Gauge(0))
	c.Set("PollCount", helpers.Counter(0))
}
