package main

import (
	"log"
	"os"

	"github.com/MarcKVR/mortgage/db"
	"github.com/MarcKVR/mortgage/handler"
	"github.com/MarcKVR/mortgage/repository"
	"github.com/MarcKVR/mortgage/service"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	_ = godotenv.Load()
	logger := db.InitLogger()

	database, err := db.GetConnection()
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close(database)

	jwtSecret := os.Getenv("JWT_SECRET")

	// Ruta raíz de la aplicación
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡CONGRATULATIONS! Welcome to te mortgage app!")
	})

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido o no proporcionado"})
		},
	})

	// Grupo de rutas protegidas bajo /admin
	paymentRepo := repository.NewRepository(database, logger)
	paymentService := service.NewService(logger, paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)
	adminGroup := app.Group("/admin", jwtMiddleware)
	adminGroup.Post("/payments", paymentHandler.Create)
	adminGroup.Get("/payments/:id", paymentHandler.Get)

	// Otro grupo de rutas protegidas bajo /api
	userRepo := repository.NewUserRepository(database, logger)
	userService := service.NewUserService(userRepo, logger)
	userHandler := handler.NewUserHandler(userService)
	apiGroup := app.Group("/api", jwtMiddleware)
	apiGroup.Post("/users", userHandler.Create)
	apiGroup.Get("/users/:id", userHandler.Get)

	// Rutas sin protección
	authRepo := repository.NewAuthRepository(database, logger)
	authService := service.NewAuthService(authRepo, logger)
	authHandler := handler.NewAuthHandler(authService)
	app.Post("/login", authHandler.Login)

	// // Agrupador de rutas
	// api := app.Group("/api")
	// v1 := api.Group("/v1")
	// v1.Get("/users", func(c *fiber.Ctx) error {
	// 	return c.SendString("Welcome to users API!")
	// })

	// // Rutas con prefijo
	// app.Route("/admin", func(api2 fiber.Router) {
	// 	api2.Get("/forms", func(c *fiber.Ctx) error {
	// 		return c.SendString("Welcome to admin FORMS Resource!")
	// 	})

	// 	api2.Get("/actions", func(c *fiber.Ctx) error {
	// 		return c.SendString("Welcome to admin ACTIONS Resource!")
	// 	})
	// })

	// Ruta por defecto de No encontrado
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Sorry can't find that page!")
	})

	log.Fatal(app.Listen(":3000"))
}
