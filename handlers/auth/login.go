package auth

import (
	"fiberJWTAuth/handlers"
	"fiberJWTAuth/search"
	"fiberJWTAuth/services"
	"fiberJWTAuth/validators"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	reqUser, err := handlers.ParseUser(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err := search.GetUser(reqUser)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if !validators.IsCorrectLoginData(reqUser, user) {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	tokenDetails, err := services.CreateJWToken(user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	return c.JSON(fiber.Map{
		"access":  tokenDetails.AccessToken,
		"refresh": tokenDetails.RefreshToken,
	})
}
