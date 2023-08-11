package repository

import (
	"clean-arch-hex/internal/domain/entity"
	"context"
)

type UserRepository interface {
	FindUser(ctx context.Context, f entity.UserFilter) (entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int64, soft bool) error
}
