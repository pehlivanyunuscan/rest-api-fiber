package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key") // JWT için gizli anahtar

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 saat geçerli olacak şekilde ayarlanıyor
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // JWT oluşturuluyor
	return token.SignedString(jwtSecret)                       // JWT imzalanıyor ve string olarak döndürülüyor
}
