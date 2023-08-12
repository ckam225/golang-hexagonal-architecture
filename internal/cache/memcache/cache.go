package memcache

import (
	"clean-arch-hex/internal/cache"
	"context"
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
func (c *Cache) Delete(ctx context.Context, key string) error {
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
func (c *Cache) DeleteAll(ctx context.Context) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = make(map[string]interface{})
	c.expiryTimes = make(map[string]time.Time)
	return nil
}

// Get implements cache.Cache.
func (c *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, found := c.data[key]
	if !found {
		return nil, fmt.Errorf("key: %s not found", key)
	}

	expiryTime, hasExpiry := c.expiryTimes[key]
	if !hasExpiry || time.Now().Before(expiryTime) {
		return value, nil
	}

	// Expired data
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
	delete(c.expiryTimes, key)
	return nil, fmt.Errorf("key: %s expired", key)
}

// Set implements cache.Cache.
func (c *Cache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
	c.expiryTimes[key] = time.Now().Add(duration)
	return nil
}

func New() cache.Cache {
	return &Cache{
		data:        make(map[string]interface{}),
		expiryTimes: make(map[string]time.Time),
	}
}
