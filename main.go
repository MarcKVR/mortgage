package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to te mortgage app!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Sorry can't find that!")
	})

	log.Fatal(app.Listen(":3000"))
}
