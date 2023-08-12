package cache

import (
	"context"
	"time"
)

//go:generate mockgen -source=cache.go -destination=mocks/mock.go
type Cache interface {
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
	DeleteAll(ctx context.Context) error
}
