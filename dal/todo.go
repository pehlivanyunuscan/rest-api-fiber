package dal

import (
	"fiber_rest/database"

	"gorm.io/gorm"
)

type Todo struct {
	ID    int
	Title string
}

func CreateTodo(todo Todo) *gorm.DB {
	return database.DB.Create(&todo)
}

func GetTodos() ([]Todo, error) {
	var todos []Todo
	res := database.DB.Find(&todos) // Veritabanından tüm Todo'ları alıyoruz
	return todos, res.Error         // Eğer hata varsa, hata döndürüyoruz
}
func GetTodo(todoID string) (Todo, error) {
	var todo Todo
	res := database.DB.First(&todo, todoID) // Veritabanından ID'ye göre Todo'yu alıyoruz
	if res.Error != nil {
		return Todo{}, res.Error // Eğer hata varsa, hata döndürüyoruz
	}
	return todo, nil // Başarılı bir şekilde alındıysa, Todo'yu döndürüyoruz
}
func UpdateTodo(todoID string, updatedTodo Todo) error {
	var todo Todo
	res := database.DB.First(&todo, todoID) // Veritabanından ID'ye göre Todo'yu alıyoruz
	if res.Error != nil {
		return res.Error // Eğer hata varsa, hata döndürüyoruz
	}

	todo.Title = updatedTodo.Title // Güncellenen başlığı alıyoruz
	res = database.DB.Save(&todo)  // Todo'yu veritabanında güncelliyoruz
	return res.Error               // Eğer hata varsa, hata döndürüyoruz
}
func DeleteTodo(todoID string) error {
	var todo Todo
	res := database.DB.First(&todo, todoID) // Veritabanından ID'ye göre Todo'yu alıyoruz
	if res.Error != nil {
		return res.Error // Eğer hata varsa, hata döndürüyoruz
	}

	res = database.DB.Delete(&todo) // Todo'yu veritabanından siliyoruz
	return res.Error                // Eğer hata varsa, hata döndürüyoruz
}
