package search

import (
	"fiberJWTAuth/config"
	"fiberJWTAuth/entities"
	"github.com/gofiber/fiber/v2"
)

func GetUser(userEntry *entities.UserReqEntry) (*entities.BaseUser, error) {
	var user entities.BaseUser

	var result = config.Database.
		Where("id = ?", userEntry.ID).
		Or("username = ?", userEntry.Username).
		Or("email = ?", userEntry.Email).
		Find(&user)

	if result.RowsAffected == 0 {
		return &user, fiber.ErrNotFound
	}

	result.Take(&user)

	return &user, nil
}
