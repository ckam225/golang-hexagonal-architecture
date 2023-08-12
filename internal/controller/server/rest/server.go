package rest

import (
	"clean-arch-hex/internal/cache"
	"clean-arch-hex/internal/controller/server"
	"clean-arch-hex/internal/db"
	"net/http"

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

// Test implements server.Server.
func (h *HTTPServer) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return h.app.Test(req, msTimeout...)
}

// Start implements server.Server.
func (h *HTTPServer) Start() error {
	fmt.Println("HTTP server is running...")
	return h.app.Listen("0.0.0.0:3000")
}

func (h HTTPServer) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "All services are running."})
}

func New(db db.Database, cache cache.Cache) server.Server {
	h := HTTPServer{
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
	h.app.Get("/posts", h.GetAllPost)
	h.app.Post("/posts", h.CreatePost)
	h.app.Get("/posts/:id<int;min(1)>", h.GetPost)
	return &h
}
