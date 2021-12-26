package handlers

import (
	"github.com/0l1v3rr/go-crud/pkg/data"
	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	book, err := data.GetBookById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	}

	return c.JSON(book)
}
