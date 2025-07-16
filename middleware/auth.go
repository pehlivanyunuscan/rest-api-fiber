package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("super-secret-key") // JWT için gizli anahtar

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization") // Authorization header'ını alıyoruz
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or invalid Authorization header",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ") // Bearer token'ı alıyoruz
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil // JWT'yi doğrulamak için gizli anahtarı kullanıyoruz
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired token",
			})
		}
		claims := token.Claims.(jwt.MapClaims) // JWT'den claim'leri alıyoruz
		c.Locals("user_id", claims["user_id"]) // Kullanıcı ID'sini context'e ekliyoruz
		return c.Next()                        // İstek işleme devam ediyor
	}
}
