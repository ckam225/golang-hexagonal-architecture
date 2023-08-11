package redis

import (
	"clean-arch-hex/internal/db"
)

type redisDB struct {
}

func New(dsn string) (db.Database, error) {
	return redisDB{}, nil
}

// Migrate implements db.Database.
func (redisDB) Migrate() error {
	panic("unimplemented")
}

// Rollback implements db.Database.
func (redisDB) Rollback() error {
	panic("unimplemented")
}
