package search

import (
	"fiberJWTAuth/config"
	"fiberJWTAuth/entities"
	"github.com/gofiber/fiber/v2"
)

type UserEntry struct {
	ID       uint
	Username string
}

func GetUser(userEntry *UserEntry) (*entities.BaseUser, error) {
	var user entities.BaseUser
	//userEntry := new(UserEntry)
	//
	//if err := c.BodyParser(userEntry); err != nil {
	//	return &user, fiber.ErrNotFound
	//}

	var result = config.Database.
		Where("id = ?", userEntry.ID).
		Or("username = ?", userEntry.Username).
		Find(&user)

	if result.RowsAffected == 0 {
		return &user, fiber.ErrNotFound
	}

	result.Take(&user)

	return &user, nil
}
