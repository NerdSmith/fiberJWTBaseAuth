package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello")
	})

	app.Listen(":8080")
}
