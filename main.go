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

	app.Post("/todos", services.CreateTodo)           // Todo oluşturma endpoint'i
	app.Get("/todos", services.GetTodos)              // Tüm Todo'ları listeleme endpoint'i
	app.Get("/todos/:todoID", services.GetTodo)       // ID'ye göre Todo alma endpoint'i
	app.Put("/todos/:todoID", services.UpdateTodo)    // ID'ye göre Todo güncelleme endpoint'i
	app.Delete("/todos/:todoID", services.DeleteTodo) // ID'ye göre Todo sil

	app.Listen("localhost:3000")
}
