package rest

import (
	"clean-arch-hex/internal/cache"
	"clean-arch-hex/internal/db"
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/usecase"
	"clean-arch-hex/internal/server"
	"context"
	"fmt"
)

type HTTPServer struct {
	db    db.Database
	cache cache.Cache
}

// Test implements server.Server.
func (h HTTPServer) Test() any {
	ucase := usecase.NewPostUseCase(h.db)
	posts, err := ucase.GetAll(context.Background(), entity.PostFilter{})
	if err != nil {
		panic(err)
	}
	return posts
}

// Start implements server.Server.
func (HTTPServer) Start() {
	fmt.Println("HTTP server is running...")
}

func New(db db.Database, cache cache.Cache) server.Server {
	return HTTPServer{
		db:    db,
		cache: cache,
	}
}
