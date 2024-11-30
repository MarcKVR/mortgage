package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	var secretKey = []byte(secret)

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims((jwt.SigningMethodHS256), claims)
	return token.SignedString(secretKey)
}
