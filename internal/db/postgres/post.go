package postgres

import (
	"clean-arch-hex/internal/domain/entity"
	"context"
	"log"

	"github.com/LineoIT/sqlb"
)

func buildFilterQuery(query string, f entity.PostFilter) *sqlb.QueryBuilder {
	if query == "" {
		query = "select * from posts"
	}
	q := sqlb.SQL(query).Limit(int64(f.Limit))

	if f.Content != "" {
		q.Where(sqlb.Ilike("content", "'%"+f.Content+"%'"))
	}
	if f.Title != "" {
		q.Where(sqlb.Ilike("title", "'%"+f.Title+"%'"))
	}
	if f.UserId > 0 {
		q.Where("user_id", f.UserId)
	}
	if f.ID > 0 {
		q.Where("id", f.ID)
	}
	return q
}

// CreatePost implements db.Database.
func (p PG) CreatePost(ctx context.Context, post *entity.Post) error {
	q := sqlb.Insert("posts").
		Value("title", post.Title).
		Value("content", post.Content).
		Return("id").
		Build()
	if err := p.db.QueryRow(ctx, q.Stmt(), q.Values()...).Scan(&post.ID); err != nil {
		return err
	}
	return nil
}

// DeletePost implements db.Database.
func (PG) DeletePost(ctx context.Context, id int64, soft bool) error {
	panic("unimplemented")
}

// FindPost implements db.Database.
func (pg PG) FindPost(ctx context.Context, f entity.PostFilter) (entity.Post, error) {
	var post entity.Post
	f.Limit = 1
	q := buildFilterQuery("", f)
	if err := pg.db.QueryRow(ctx, q.Stmt(), q.Args()...).
		Scan(&post.ID, &post.Title, &post.Content, &post.UserId); err != nil {
		log.Println(q.Debug())
		return entity.Post{}, err
	}
	return post, nil
}

// GetPosts implements db.Database.
func (pg PG) GetPosts(ctx context.Context, f entity.PostFilter) ([]entity.Post, error) {
	q := buildFilterQuery("", f)
	rows, err := pg.db.Query(ctx, q.Stmt(), q.Args()...)
	if err != nil {
		log.Println(q.Debug())
		return nil, err
	}
	defer rows.Close()

	posts := make([]entity.Post, 0)
	for rows.Next() {
		var post entity.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserId); err != nil {
			log.Println(q.Debug())
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
