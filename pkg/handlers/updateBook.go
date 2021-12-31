package handlers

import (
	"fmt"

	"github.com/0l1v3rr/go-crud/pkg/data"
	"github.com/gofiber/fiber/v2"
)

func UpdateBook(c *fiber.Ctx) error {
	book, err := data.GetBookById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	}

	for index, item := range data.Books {
		if item.ID == book.ID {
			book := new(data.Book)
			if err := c.BodyParser(book); err != nil {
				fmt.Println(err)
				return c.SendStatus(200)
			}
			data.Books[index] = *book
			break
		}
	}

	return c.JSON(data.Books)
}
