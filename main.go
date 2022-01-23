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
	err := config.Connect()
	if err != nil {
		panic(err)
	}

	app.Post("/signup", auth.Signup)
	app.Post("/login", auth.Login)
	app.Post("/refresh", auth.Refresh)

	app.Listen(":8080")
}
