package repository

import (
	"clean-arch-hex/internal/domain/entity"
	"context"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type PostRepository interface {
	GetPosts(ctx context.Context, f entity.PostFilter) ([]entity.Post, error)
	FindPost(ctx context.Context, f entity.PostFilter) (entity.Post, error)
	CreatePost(ctx context.Context, post *entity.Post) error
	UpdatePost(ctx context.Context, post *entity.Post) error
	DeletePost(ctx context.Context, id int64, soft bool) error
}

type UserRepository interface {
	FindUser(ctx context.Context, f entity.UserFilter) (entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int64, soft bool) error
}
