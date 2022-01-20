package main

import (
	"fiberJWTAuth/search"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

//type MyCustomClaims struct {
//	UID uint64 `json:"UID"`
//	jwt.RegisteredClaims
//}

type LoginRequestBody struct {
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	loginBody := new(LoginRequestBody)
	var userEntry = new(search.UserEntry)
	c.BodyParser(userEntry)
	fmt.Println(search.GetUser(userEntry))

	if err := c.BodyParser(loginBody); err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	//user := c.FormValue("username")
	//pass := c.FormValue("password")

	// Throws Unauthorized error
	if loginBody.Username != TUser.Username || loginBody.Password != TUser.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	//Create the Claims
	tokenClaims := jwt.MapClaims{
		"UID":      1,
		"username": "username",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"UID":      1,
		"username": "username",
		"exp":      time.Now().Add(time.Hour * 24 * 10).Unix(),
	}

	//tokenClaims := &jwt.RegisteredClaims{
	//	ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
	//}

	//atClaims := jwt.MapClaims{}
	//atClaims["authorized"] = true
	//atClaims["access_uuid"] = uuid.NewV4().String()
	//atClaims["user_id"] = TUser.UID
	//atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	t, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	r, err := refresh.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"access":  t,
		"refresh": r,
	})
}
