package oracle

import (
	"clean-arch-hex/internal/db"
	"database/sql"
)

type Oracle struct {
	db *sql.DB
}

func New(dsn string) (db.Database, error) {
	openDB, err := sql.Open("oracle", dsn)
	if err != nil {
		return Oracle{}, err
	}
	return Oracle{
		db: openDB,
	}, nil
}

// Migrate implements db.Database.
func (Oracle) Migrate() error {
	panic("unimplemented")
}

// Rollback implements db.Database.
func (Oracle) Rollback() error {
	panic("unimplemented")
}
