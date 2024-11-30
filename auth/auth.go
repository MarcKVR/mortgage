package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret")

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token válido por 72 horas
	}

	token := jwt.NewWithClaims((jwt.SigningMethodHS256), claims)
	return token.SignedString(jwtSecret)
}
