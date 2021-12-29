package handlers

import (
	"log"

	"github.com/0l1v3rr/go-crud/pkg/data"
	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	book := new(data.Book)

	if err := c.BodyParser(book); err != nil {
		log.Fatalf("%v", err)
		return c.SendStatus(500)
	}

	data.Books = append(data.Books, *book)
	return c.SendStatus(fiber.StatusCreated)
}
