package rest

import (
	"clean-arch-hex/internal/cache"
	"clean-arch-hex/internal/db"
	"clean-arch-hex/internal/server"
	"clean-arch-hex/internal/server/rest/handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	db    db.Database
	cache cache.Cache
	app   *fiber.App
}

// Start implements server.Server.
func (h *HTTPServer) Start() error {
	h.app.Get("/posts", h.GetAllPost)
	h.app.Get("/posts/:id<int;min(1)>", h.GetPost)

	h.app.Mount("/articles", handler.Posts(h.db).Endpoints())
	fmt.Println("HTTP server is running...")
	return h.app.Listen("0.0.0.0:3000")

}

func New(db db.Database, cache cache.Cache) server.Server {
	return &HTTPServer{
		db:    db,
		cache: cache,
		app:   fiber.New(),
	}
}
