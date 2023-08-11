package rest

import (
	"clean-arch-hex/internal/cache"
	"clean-arch-hex/internal/controller/server"
	"clean-arch-hex/internal/db"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HTTPServer struct {
	db    db.Database
	cache cache.Cache
	app   *fiber.App
}

// Start implements server.Server.
func (h *HTTPServer) Start() error {

	h.app.Use(logger.New())
	h.app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT",
		AllowCredentials: true,
	}))
	//h.app.Static("/static", filepath.Join("/", "public"))

	// posts endpoints
	h.app.Get("/posts", h.GetAllPost)
	h.app.Get("/posts/:id<int;min(1)>", h.GetPost)

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
