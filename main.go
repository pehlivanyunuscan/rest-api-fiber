package main

import (
	"fiber_rest/dal"
	"fiber_rest/database"
	"fiber_rest/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()                   // Başta bir kez bağlantı kuruyoruz
	database.DB.AutoMigrate(&dal.Todo{}) // Veritabanı otomatik olarak Todo modeline göre güncelleniyor

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{ // Ana sayfada basit bir JSON mesajı döndürüyoruz
			"message": "Welcome to the Fiber REST API!",
		})
	})

	app.Post("/todos", services.CreateTodo) // Todo oluşturma endpoint'i
	app.Get("/todos", services.GetTodos)    // Tüm Todo'ları listeleme endpoint'i

	app.Listen("localhost:3000")
}
