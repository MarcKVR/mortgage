package router

import (
	"os"

	"github.com/MarcKVR/mortgage/handler"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	jwtSecret := os.Getenv("JWT_SECRET")

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido o no proporcionado"})
		},
	})

	apiGroup := app.Group("/api", jwtMiddleware)
	apiGroup.Post("/users", userHandler.Create)
	apiGroup.Get("/users/:id", userHandler.Get)
	apiGroup.Get("/users", userHandler.GetUsers)
	apiGroup.Put("/users/:id", userHandler.Update)
}
