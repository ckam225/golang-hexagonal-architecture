package handler

import (
	"clean-arch-hex/internal/db"
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/usecase"
	"context"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	*fiber.App
	db db.Database
}

func Posts(db db.Database) *PostHandler {
	return &PostHandler{
		App: fiber.New(),
		db:  db,
	}
}

func (r *PostHandler) Endpoints() *fiber.App {
	r.Get("/posts", r.GetAllPost)
	r.Get("/posts/:id<int;min(1)>", r.GetPost)
	return r.App
}

func (h *PostHandler) GetAllPost(c *fiber.Ctx) error {
	ucase := usecase.ReadPost(h.db)
	posts, err := ucase.GetAll(context.Background(), entity.PostFilter{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return c.JSON(posts)
}

func (h PostHandler) GetPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}
	ucase := usecase.ReadPost(h.db)
	posts, err := ucase.Find(context.Background(), entity.PostFilter{
		ID: id,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return c.JSON(posts)
}
