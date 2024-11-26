package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Mortgage control",
	})

	// Ruta raíz de la aplicación
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to te mortgage app!")
	})

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

	dsn := "host=localhost user=mortgage_user password=mortgage_pass dbname=mortgage_db port=5439 sslmode=disable TimeZone=America/Mexico_City"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection to the database was successful!")
	}

	log.Fatal(app.Listen(":3000"))
}
