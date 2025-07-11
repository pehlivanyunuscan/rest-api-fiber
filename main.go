package main

import (
	"fiber_rest/database"

	"github.com/gofiber/fiber/v3"
)

func main() {

	database.Connect() // Başta bir kez bağlantı kuruyoruz

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen("localhost:3000")
}
