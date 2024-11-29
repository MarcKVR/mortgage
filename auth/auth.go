package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret")

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token válido por 72 horas
	}

	token := jwt.NewWithClaims((jwt.SigningMethodHS256), claims)
	return token.SignedString(jwtSecret)
}
