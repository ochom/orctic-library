package utils

// Cache ...
type Cache struct {
	data map[string]interface{}
}

// NewCache ...
func NewCache() *Cache {
	return &Cache{data: make(map[string]interface{})}
}

// Set ...
func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

// Get ...
func (c *Cache) Get(key string) interface{} {
	return c.data[key]
}

// Has ...
func (c *Cache) Has(key string) bool {
	_, ok := c.data[key]
	return ok
}

// Delete ...
func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

// Clear ...
func (c *Cache) Clear() {
	c.data = make(map[string]interface{})
}
