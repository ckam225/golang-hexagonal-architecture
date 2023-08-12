package db

import (
	"clean-arch-hex/internal/domain/repository"
)

//go:generate mockgen -source=db.go -destination=mocks/mock.go
type Database interface {
	repository.UserRepository
	repository.PostRepository
	Migrate() error
	Rollback() error
}
