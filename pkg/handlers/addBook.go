package handlers

import (
	"fmt"

	"github.com/0l1v3rr/go-crud/pkg/data"
	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	book := new(data.Book)

	if err := c.BodyParser(book); err != nil {
		fmt.Println(err)
		return c.SendStatus(200)
	}

	data.Books = append(data.Books, *book)
	return c.SendStatus(200)
}
