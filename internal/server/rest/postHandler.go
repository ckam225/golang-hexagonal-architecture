package rest

import (
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/usecase"
	"context"

	"github.com/gofiber/fiber/v2"
)

func (h HTTPServer) GetAllPost(c *fiber.Ctx) error {
	ucase := usecase.ReadPost(h.db)
	posts, err := ucase.GetAll(context.Background(), entity.PostFilter{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return c.JSON(posts)
}

func (h HTTPServer) GetPost(c *fiber.Ctx) error {
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
