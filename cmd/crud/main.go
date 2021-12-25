package main

import (
	"fmt"

	"github.com/0l1v3rr/go-crud/pkg/data"
	"github.com/0l1v3rr/go-crud/pkg/handlers"
	"github.com/gofiber/fiber/v2"
)

const port = 8080

func main() {
	data.CreateBooks()
	app := fiber.New()

	app.Get("/api/books", handlers.GetBooks)

	app.Listen(fmt.Sprintf(":%v", port))
}
