package redis

import (
	"clean-arch-hex/internal/domain/entity"
	"context"
)

// CreatePost implements db.Database.
func (redisDB) CreatePost(ctx context.Context, post *entity.Post) error {
	panic("unimplemented")
}

// DeletePost implements db.Database.
func (redisDB) DeletePost(ctx context.Context, id int64, soft bool) error {
	panic("unimplemented")
}

// FindPost implements db.Database.
func (redisDB) FindPost(ctx context.Context, f entity.PostFilter) (entity.Post, error) {
	panic("unimplemented")
}

// GetPosts implements db.Database.
func (redisDB) GetPosts(ctx context.Context, f entity.PostFilter) ([]entity.Post, error) {
	return []entity.Post{
		{
			ID:      1,
			Title:   "Java for null",
			Content: "Course for JAVA",
		},
		{
			ID:      2,
			Title:   "Net/HTTP",
			Content: "Golang HTTP module",
		},
	}, nil
}

// UpdatePost implements db.Database.
func (redisDB) UpdatePost(ctx context.Context, post *entity.Post) error {
	panic("unimplemented")
}
