package handlers

import (
	"github.com/0l1v3rr/go-crud/pkg/data"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(data.Books)
}
