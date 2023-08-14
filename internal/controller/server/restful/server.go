package restful

import (
	"clean-arch-hex/internal/cache"
	srv "clean-arch-hex/internal/controller/server"
	"clean-arch-hex/internal/db"
	"net/http"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type server struct {
	db    db.Database
	cache cache.Cache
	app   *fiber.App
}

// Test implements server.Server.
func (h *server) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return h.app.Test(req, msTimeout...)
}

// Start implements server.Server.
func (h *server) Start(address string) error {
	fmt.Println("HTTP server is running...")
	return h.app.Listen(address)
}

func (h server) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "All services are running."})
}

func New(db db.Database, cache cache.Cache) srv.Server {
	h := server{
		db:    db,
		cache: cache,
		app:   fiber.New(),
	}
	h.app.Use(logger.New())
	h.app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT",
		AllowCredentials: true,
	}))
	// h.app.Static("/static", filepath.Join("/", "public"))

	// posts endpoints
	h.app.Get("/health", h.Health)
	h.app.Get("/posts", h.GetAllPosts)
	h.app.Post("/posts", h.CreatePost)
	h.app.Get("/posts/:id<int;min(1)>", h.GetPost)
	return &h
}
