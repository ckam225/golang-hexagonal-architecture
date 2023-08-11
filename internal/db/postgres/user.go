package postgres

import (
	"clean-arch-hex/internal/domain/entity"
	"context"
)

// CreateUser implements db.Database.
func (PG) CreateUser(ctx context.Context, user *entity.User) error {
	panic("unimplemented")
}

// DeleteUser implements db.Database.
func (PG) DeleteUser(ctx context.Context, id int64, soft bool) error {
	panic("unimplemented")
}

// FindUser implements db.Database.
func (PG) FindUser(ctx context.Context, f entity.UserFilter) (entity.User, error) {
	panic("unimplemented")
}

// UpdateUser implements db.Database.
func (PG) UpdateUser(ctx context.Context, user *entity.User) error {
	panic("unimplemented")
}
