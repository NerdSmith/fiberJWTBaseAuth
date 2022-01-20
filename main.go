package main

import (
	"fiberJWTAuth/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	config.Connect()

	app.Post("/login", Login)
	//app.Post("/login", Login)
	//app.Post("/login", Login)

	app.Listen(":8080")

	//config.Connect()
}
