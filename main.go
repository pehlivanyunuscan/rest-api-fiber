package main

import (
	"fiber_rest/dal"
	"fiber_rest/database"

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

	type TodoCreate struct {
		Title string // Todo oluşturma için gerekli alan
	}
	app.Post("/todos", func(c *fiber.Ctx) error {
		t := new(TodoCreate) // TodoCreate tipinde yeni bir değişken oluşturuyoruz
		if err := c.BodyParser(t); err != nil {
			return err
		}

		newTodo := dal.Todo{
			Title: t.Title, // Kullanıcıdan gelen başlığı alıyoruz
		}
		res := database.DB.Create(&newTodo) // Veritabanına yeni Todo ekliyoruz

		if res.Error != nil {
			return res.Error // Eğer hata varsa, hatayı döndürüyoruz
		}

		return c.JSON(fiber.Map{ // Başarılı bir şekilde eklenirse, yeni Todo'yu JSON olarak döndürüyoruz
			"message": "Todo created successfully!",
		})
	})

	app.Listen("localhost:3000")
}
