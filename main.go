package main

import (
	"fiber_rest/dal"
	"fiber_rest/database"
	"fiber_rest/middleware"
	"fiber_rest/services"
	"fiber_rest/types"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	database.Connect()                     // Başta bir kez bağlantı kuruyoruz
	database.DB.AutoMigrate(&dal.Todo{})   // Veritabanı otomatik olarak Todo modeline göre güncelleniyor
	database.DB.AutoMigrate(&types.User{}) // Veritabanı otomatik olarak User modeline göre güncelleniyor
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

	app.Post("/register", func(c *fiber.Ctx) error {
		var input types.User
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid input",
			})
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
		input.Password = string(hashedPassword)
		if err := dal.CreateUser(database.DB, &input); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create user",
			})
		}
		return c.JSON(fiber.Map{
			"message": "User created successfully",
			"user":    input,
		})
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		var input types.User
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid input",
			})
		}
		user, err := dal.GetUserByEmail(database.DB, input.Email)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid email or password",
			})
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid email or password",
			})
		}
		token, err := services.GenerateJWT(user.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to generate token",
			})
		}
		return c.JSON(fiber.Map{
			"message": "Login successful",
			"token":   token,
		})
	})
	app.Get("/profile", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")                                 // JWT'den kullanıcı ID'sini alıyoruz
		user, err := dal.GetUserByEmail(database.DB, userID.(string)) // Kullanıcıyı veritabanından alıyoruz
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.JSON(user) // Kullanıcı bilgilerini JSON olarak döndürüyoruz
	})

	app.Listen("localhost:3000")
}
