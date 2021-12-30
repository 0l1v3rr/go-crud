package handlers

import (
	"github.com/0l1v3rr/go-crud/pkg/data"
	"github.com/gofiber/fiber/v2"
)

func DeleteBook(c *fiber.Ctx) error {
	book, err := data.GetBookById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	}

	for index, item := range data.Books {
		if item.ID == book.ID {
			data.Books = append(data.Books[:index], data.Books[index+1:]...)
			break
		}
	}

	return c.JSON(data.Books)
}
