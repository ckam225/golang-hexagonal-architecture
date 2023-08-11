package postgres

import (
	"clean-arch-hex/internal/domain/entity"
	"context"
	"fmt"
	"log"
)

func buildFilterQuery(query string, f entity.PostFilter) (string, []interface{}) {
	if query == "" {
		query = "select * from posts"
	}
	var clause string
	args := make([]interface{}, 0)
	if f.Content != "" {
		clause += "content ilike '%$1%'"
		args = append(args, f.Content)
	}
	if f.Title != "" {
		if clause != "" {
			clause += " and "
		}
		clause += fmt.Sprintf("title=$%d", len(args)+1)
		args = append(args, f.Title)
	}
	if f.UserId > 0 {
		if clause != "" {
			clause += " and "
		}
		clause += fmt.Sprintf("user_id=$%d", len(args)+1)
		args = append(args, f.UserId)
	}
	if f.ID > 0 {
		if clause != "" {
			clause += " and "
		}
		clause += fmt.Sprintf("id=$%d", len(args)+1)
		args = append(args, f.ID)
	}
	if clause != "" {
		query += " where " + clause
	}
	return query, args
}

// CreatePost implements db.Database.
func (PG) CreatePost(ctx context.Context, post *entity.Post) error {
	panic("unimplemented")
}

// DeletePost implements db.Database.
func (PG) DeletePost(ctx context.Context, id int64, soft bool) error {
	panic("unimplemented")
}

// FindPost implements db.Database.
func (pg PG) FindPost(ctx context.Context, f entity.PostFilter) (entity.Post, error) {
	var post entity.Post
	q, args := buildFilterQuery("select * from posts limit 1", f)
	if err := pg.db.QueryRow(ctx, q, args...).
		Scan(&post.ID, &post.Title, &post.Content, &post.UserId); err != nil {
		return entity.Post{}, err
	}
	return post, nil
}

// GetPosts implements db.Database.
func (pg PG) GetPosts(ctx context.Context, f entity.PostFilter) ([]entity.Post, error) {
	posts := make([]entity.Post, 0)
	q, args := buildFilterQuery("select * from posts", f)
	rows, err := pg.db.Query(ctx, q, args...)
	if err != nil {
		log.Println(q, args)
		return nil, err
	}
	for rows.Next() {
		var post entity.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserId); err != nil {
			log.Println(q, args)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// UpdatePost implements db.Database.
func (PG) UpdatePost(ctx context.Context, post *entity.Post) error {
	panic("unimplemented")
}
