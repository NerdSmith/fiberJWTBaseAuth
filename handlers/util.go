package handlers

import (
	"fiberJWTAuth/entities"
	"github.com/gofiber/fiber/v2"
)

func ParseUser(c *fiber.Ctx) (*entities.UserReqEntry, error) {
	var user = new(entities.UserReqEntry)
	if err := c.BodyParser(user); err != nil {
		return user, err
	}
	return user, nil
}
