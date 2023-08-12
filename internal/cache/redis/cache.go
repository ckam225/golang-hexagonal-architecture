package redis

import (
	"clean-arch-hex/internal/cache"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

// Delete implements cache.Cache.
func (c *Cache) Delete(ctx context.Context, key string) error {
	panic("unimplemented")
}

// DeleteAll implements cache.Cache.
func (*Cache) DeleteAll(ctx context.Context) error {
	panic("unimplemented")
}

// Get implements cache.Cache.
func (c *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	return c.client.Get(ctx, "key").Result()
}

// Set implements cache.Cache.
func (c *Cache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return c.client.Set(ctx, key, value, duration).Err()
}

func NewCache() cache.Cache {
	url := "redis://localhost:6379?password=hello&protocol=3"
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	return &Cache{
		client: redis.NewClient(opts),
	}
}
