package main

import (
	"log"

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

	// Ruta raíz de la aplicación
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡CONGRATULATIONS! Welcome to te mortgage app!")
	})

	userRepo := repository.NewUserRepository(database, logger)
	userService := service.NewUserService(userRepo, logger)
	userHandler := handler.NewUserHandler(userService)
	app.Post("/users", userHandler.Create)
	app.Get("/users/:id", userHandler.Get)

	paymentRepo := repository.NewRepository(database, logger)
	paymentService := service.NewService(logger, paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	app.Use("/payments", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{"error": "Unauthorized for resource"})
		},
	}))

	// app.Route("/admin", func(adminApi fiber.Router) {
	// 	adminApi.Post("/payments", paymentHandler.Create)
	// 	adminApi.Get("/payments/:id", paymentHandler.Get)

	// })

	app.Post("/payments", paymentHandler.Create)
	app.Get("/payments/:id", paymentHandler.Get)

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

	// // Ruta por defecto de No encontrado
	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.Status(404).SendString("Sorry can't find that!")
	// })

	// dsn := "host=localhost user=mortgage_user password=mortgage_pass dbname=mortgage_db port=5439 sslmode=disable TimeZone=America/Mexico_City"
	// _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Connection to the database was successful!")
	// }

	log.Fatal(app.Listen(":3000"))
}
