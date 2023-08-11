package repository

import (
	"clean-arch-hex/internal/domain/entity"
	"context"
)

type PostRepository interface {
	GetPosts(ctx context.Context, f entity.PostFilter) ([]entity.Post, error)
	FindPost(ctx context.Context, f entity.PostFilter) (entity.Post, error)
	CreatePost(ctx context.Context, post *entity.Post) error
	UpdatePost(ctx context.Context, post *entity.Post) error
	DeletePost(ctx context.Context, id int64, soft bool) error
}
