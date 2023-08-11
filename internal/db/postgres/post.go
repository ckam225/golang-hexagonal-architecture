package postgres

import (
	"clean-arch-hex/internal/domain/entity"
	"context"
)

// CreatePost implements db.Database.
func (PG) CreatePost(ctx context.Context, post *entity.Post) error {
	panic("unimplemented")
}

// DeletePost implements db.Database.
func (PG) DeletePost(ctx context.Context, id int64, soft bool) error {
	panic("unimplemented")
}

// FindPost implements db.Database.
func (PG) FindPost(ctx context.Context, f entity.PostFilter) (entity.Post, error) {
	panic("unimplemented")
}

// GetPosts implements db.Database.
func (pg PG) GetPosts(ctx context.Context, f entity.PostFilter) ([]entity.Post, error) {
	return []entity.Post{
		{
			ID:      1,
			Title:   "Java for null",
			Content: "Course for JAVA",
			UserId:  0,
		},
		{
			ID:      2,
			Title:   "Net/HTTP",
			Content: "Golang HTTP module",
			UserId:  0,
		},
	}, nil
}

// UpdatePost implements db.Database.
func (PG) UpdatePost(ctx context.Context, post *entity.Post) error {
	panic("unimplemented")
}
