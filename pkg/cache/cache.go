package cache

import "sync"

// InMemoryCache implements the CacheManager interface using a map
type InMemoryCache struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewInMemoryCache initializes a new InMemoryCache instance
func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]string),
	}
}

// Get retrieves a value from the cache
func (c *InMemoryCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

// Set adds or updates a value in the cache
func (c *InMemoryCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Flush clears all entries in the cache
func (c *InMemoryCache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]string)
}
