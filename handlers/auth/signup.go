package auth

import (
	"fiberJWTAuth/config"
	"fiberJWTAuth/entities"
	"fiberJWTAuth/handlers"
	"fiberJWTAuth/search"
	"fiberJWTAuth/validators"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func isExist(user *entities.UserReqEntry) bool {
	_, err := search.GetUser(user)
	if err != nil {
		return false
	}
	return true
}

func verifyReqUser(reqUser *entities.UserReqEntry) bool {
	return validators.VerifyUsername(reqUser) &&
		validators.VerifyPassword(reqUser) &&
		validators.VerifyEmail(reqUser) &&
		(!isExist(reqUser))
}

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
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	fmt.Println(reqUser)
	if !(verifyReqUser(reqUser)) {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err := createNewUser(reqUser)
	return c.Status(201).JSON(user)
}
