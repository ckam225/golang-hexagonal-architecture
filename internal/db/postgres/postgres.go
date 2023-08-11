package postgres

import (
	"clean-arch-hex/internal/db"
	"clean-arch-hex/pkg/utils"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	// _ "github.com/lib/pq"
)

type PG struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, dsn string, maxAttempts int) (db.Database, error) {
	var pool *pgxpool.Pool
	var err error
	if err = utils.TryAttempt(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second); err != nil {
		return nil, errors.New("error do with tries postgresql")
	}
	return PG{
		db: pool,
	}, nil
}

// Migrate implements db.Database.
func (PG) Migrate() error {
	panic("unimplemented")
}

// Rollback implements db.Database.
func (PG) Rollback() error {
	panic("unimplemented")
}
