package memcache

import (
	"clean-arch-hex/internal/cache"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data        map[string]interface{}
	expiryTimes map[string]time.Time
	mutex       sync.RWMutex
}

// Delete implements cache.Cache.
func (c *Cache) Delete(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, found := c.data[key]
	if !found {
		return fmt.Errorf("key %s not found in cache", key)
	}
	delete(c.data, key)
	delete(c.expiryTimes, key)
	return nil
}

// DeleteAll implements cache.Cache.
func (c *Cache) DeleteAll() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = make(map[string]interface{})
	c.expiryTimes = make(map[string]time.Time)
}

// Get implements cache.Cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, found := c.data[key]
	if !found {
		return nil, false
	}

	expiryTime, hasExpiry := c.expiryTimes[key]
	if !hasExpiry || time.Now().Before(expiryTime) {
		return value, true
	}

	// Expired data
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
	delete(c.expiryTimes, key)
	return nil, false
}

// Set implements cache.Cache.
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
	c.expiryTimes[key] = time.Now().Add(duration)
}

func New() cache.Cache {
	return &Cache{
		data:        make(map[string]interface{}),
		expiryTimes: make(map[string]time.Time),
	}
}
