package utils

import "sync"

// Cache ...
type Cache struct {
	data map[string]any
	mu   sync.RWMutex
}

// NewCache ...
func NewCache() *Cache {
	return &Cache{data: make(map[string]any)}
}

// Set ...
func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

// Get ...
func (c *Cache) Get(key string) any {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data[key]
}

// Has ...
func (c *Cache) Has(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.data[key]
	return ok
}

// Delete ...
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

// Clear ...
func (c *Cache) Clear() {
	c.mu.Lock()
	c.data = make(map[string]any)
	c.mu.Unlock()
}
