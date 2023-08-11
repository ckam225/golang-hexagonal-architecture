package db

import (
	"clean-arch-hex/internal/domain/repository"
)

type Database interface {
	repository.UserRepository
	repository.PostRepository
	Migrate() error
	Rollback() error
}
