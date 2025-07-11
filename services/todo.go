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
func GetTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID") // URL parametresinden Todo ID'sini alıyoruz
	if todoID == "" {
		return c.Status(400).JSON(fiber.Map{ // Eğer ID boş ise, 400 hatası döndürüyoruz
			"message": "Todo ID is required",
		})
	}

	todo, err := dal.GetTodo(todoID) // Veritabanından ID'ye göre Todo'yu alıyoruz
	if err != nil {
		return c.Status(404).JSON(fiber.Map{ // Eğer Todo bulunamazsa, 404 hatası döndürüyoruz
			"message": "Todo not found",
		})
	}

	return c.JSON(todo) // Başarılı bir şekilde alındıysa, Todo'yu JSON olarak döndürüyoruz
}
func UpdateTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID") // URL parametresinden Todo ID'sini alıyoruz
	if todoID == "" {
		return c.Status(400).JSON(fiber.Map{ // Eğer ID boş ise, 400 hatası döndürüyoruz
			"message": "Todo ID is required",
		})
	}

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

	todo, err := dal.GetTodo(todoID) // Veritabanından ID'ye göre Todo'yu alıyoruz
	if err != nil {
		return c.Status(404).JSON(fiber.Map{ // Eğer Todo bulunamazsa, 404 hatası döndürüyoruz
			"message": "Todo not found",
		})
	}

	todo.Title = t.Title               // Kullanıcıdan gelen başlığı güncelliyoruz
	err = dal.UpdateTodo(todoID, todo) // Güncellenmiş Todo'yu veritabanına kaydediyoruz
	if err != nil {
		return c.Status(500).JSON(fiber.Map{ // Eğer veritabanına kaydetme işlemi başarısız olursa, 500 hatası döndürüyoruz
			"message": "Failed to update Todo",
		})
	}

	return c.JSON(fiber.Map{ // Başarılı bir şekilde güncellendiyse, güncellenmiş Todo'yu JSON olarak döndürüyoruz
		"message": "Todo updated successfully!",
	})
}
func DeleteTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID") // URL parametresinden Todo ID'sini alıyoruz
	if todoID == "" {
		return c.Status(400).JSON(fiber.Map{ // Eğer ID boş ise, 400 hatası döndürüyoruz
			"message": "Todo ID is required",
		})
	}

	err := dal.DeleteTodo(todoID) // Veritabanından ID'ye göre Todo'yu siliyoruz
	if err != nil {
		return c.Status(404).JSON(fiber.Map{ // Eğer Todo bulunamazsa, 404 hatası döndürüyoruz
			"message": "Todo not found",
		})
	}

	return c.JSON(fiber.Map{ // Başarılı bir şekilde silindiyse, başarı mesajı döndürüyoruz
		"message": "Todo deleted successfully!",
	})
}
