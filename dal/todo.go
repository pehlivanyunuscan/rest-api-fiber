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
