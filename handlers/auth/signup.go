package auth

import (
	"fiberJWTAuth/config"
	"fiberJWTAuth/entities"
	"fiberJWTAuth/handlers"
	"fiberJWTAuth/validators"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func createNewUser(reqUser *entities.UserReqEntry) (*entities.BaseUser, error) {
	var user = new(entities.BaseUser)
	user.Username = reqUser.Username
	user.Password = reqUser.Password
	user.Email = reqUser.Email

	config.Database.Create(&user)
	return user, nil
}

func Signup(c *fiber.Ctx) error {
	reqUser, err := handlers.ParseUser(c)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	fmt.Println(reqUser)
	if !(validators.VerifyReqUser(reqUser)) {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err := createNewUser(reqUser)
	return c.Status(201).JSON(user)
}
