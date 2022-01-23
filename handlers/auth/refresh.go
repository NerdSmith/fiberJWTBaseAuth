package auth

import (
	"fiberJWTAuth/entities"
	"fiberJWTAuth/search"
	"fiberJWTAuth/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
)

func parseRefreshToken(reqRefreshToken string) (*jwt.Token, error) {
	rToken, err := jwt.Parse(reqRefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	return rToken, err
}

func createNewAccessTokenByUserID(userID uint) (string, error) {
	reqUser := new(entities.UserReqEntry)
	JWTokenDetails := new(entities.JWTokenDetails)
	reqUser.ID = userID
	user, err := search.GetUser(reqUser)
	if err != nil {
		return "", err
	}
	if err := services.CreateAccessToken(user, JWTokenDetails); err != nil {
		return "", err
	}
	return JWTokenDetails.AccessToken, nil
}

func Refresh(c *fiber.Ctx) error {
	refreshToken := new(
		struct {
			Refresh string
		})
	if err := c.BodyParser(&refreshToken); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	rToken, err := parseRefreshToken(refreshToken.Refresh)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	rTokenClaims, ok := rToken.Claims.(jwt.MapClaims)
	if !ok || !rToken.Valid {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	userID, err := strconv.ParseUint(fmt.Sprintf("%.f", rTokenClaims["ID"]), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	newAccessToken, err := createNewAccessTokenByUserID(uint(userID))
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	return c.JSON(fiber.Map{
		"access": newAccessToken,
	})
}
