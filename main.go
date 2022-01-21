package main

import (
	"fiberJWTAuth/config"
	"fiberJWTAuth/handlers/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	config.Connect()

	app.Post("/login", Login)
	app.Post("/signup", auth.Signup)
	//app.Post("/login", Login)
	//app.Post("/login", Login)

	app.Listen(":8080")

	//config.Connect()
}
