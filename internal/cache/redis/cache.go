package redis

import (
	"clean-arch-hex/internal/cache"
	"time"
)

type Cache struct {
}

// Delete implements cache.Cache.
func (*Cache) Delete(key string) error {
	panic("unimplemented")
}

// DeleteAll implements cache.Cache.
func (*Cache) DeleteAll() {
	panic("unimplemented")
}

// Get implements cache.Cache.
func (*Cache) Get(key string) (interface{}, bool) {
	panic("unimplemented")
}

// Set implements cache.Cache.
func (*Cache) Set(key string, value interface{}, duration time.Duration) {
	panic("unimplemented")
}

func NewCache() cache.Cache {
	return &Cache{}
}
