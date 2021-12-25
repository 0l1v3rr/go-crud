package main

import (
	"fmt"

	"github.com/0l1v3rr/go-crud/pkg/data"
	"github.com/0l1v3rr/go-crud/pkg/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const port = 8080

func main() {
	data.CreateBooks()
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	app.Get("/api/books", handlers.GetBooks)

	app.Listen(fmt.Sprintf(":%v", port))
}
