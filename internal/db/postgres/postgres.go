package postgres

import (
	"clean-arch-hex/internal/db"
	"database/sql"
	// _ "github.com/lib/pq"
)

type PG struct {
	db *sql.DB
}

func New(dsn string) (db.Database, error) {
	// openDB, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	return PG{}, err
	// }
	return PG{
		// db: openDB,
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
