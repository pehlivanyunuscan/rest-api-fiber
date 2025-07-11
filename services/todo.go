package services

import (
	"fiber_rest/dal"
	"fiber_rest/types"

	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	t := new(types.TodoCreate) // TodoCreate tipinde yeni bir değişken oluşturuyoruz

	if err := c.BodyParser(t); err != nil {
		return c.Status(400).JSON(fiber.Map{ // Eğer body parse edilemezse, 400 hatası döndürüyoruz
			"message": "BAD REQUEST",
		})
	}

	if t.Title == "" { // Eğer başlık boş ise, 400 hatası döndürüyoruz
		return c.Status(400).JSON(fiber.Map{
			"message": "Title is required",
		})
	}

	newTodo := dal.Todo{
		Title: t.Title, // Kullanıcıdan gelen başlığı alıyoruz
	}
	res := dal.CreateTodo(newTodo) // Yeni Todo'yu veritabanına ekliyoruz
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{ // Eğer veritabanına ekleme işlemi başarısız olursa, 500 hatası döndürüyoruz
			"message": "Failed to create Todo",
		})
	}

	return c.JSON(fiber.Map{ // Başarılı bir şekilde eklenirse, yeni Todo'yu JSON olarak döndürüyoruz
		"message": "Todo created successfully!",
	})
}

func GetTodos(c *fiber.Ctx) error {
	todos, err := dal.GetTodos() // Veritabanından tüm Todo'ları alıyoruz
	if err != nil {
		return c.Status(500).JSON(fiber.Map{ // Eğer hata varsa, 500 hatası döndürüyoruz
			"message": "Failed to retrieve Todos",
		})
	}

	return c.JSON(todos) // Başarılı bir şekilde alındıysa, Todo'ları JSON olarak döndürüyoruz

}
