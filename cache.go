package cache

import (
	"sync"
)

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

// New создает новый экземпляр Cache
func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set сохраняет значение в кэш
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get получает значение из кэша
func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// Delete удаляет значение из кэша
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
