package services

import (
	"errors"
	"fiberJWTAuth/config"
	"fiberJWTAuth/entities"
	"github.com/golang-jwt/jwt/v4"
	"github.com/twinj/uuid"
	"os"
	"time"
)

func CreateAccessToken(user *entities.BaseUser, tokenDetails *entities.JWTokenDetails) error {
	tokenDetails.AtExpires = time.Now().Add(config.ACCESS_EXP_TIME).Unix()
	tokenDetails.AccessUuid = uuid.NewV4().String()

	accessTokenClaims := jwt.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"exp":      tokenDetails.AtExpires,
	}

	var err error

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	tokenDetails.AccessToken, err = access.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return errors.New("can't create access token")
	}

	return nil
}

func CreateRefreshToken(user *entities.BaseUser, tokenDetails *entities.JWTokenDetails) error {
	tokenDetails.RtExpires = time.Now().Add(config.REFRESH_EXP_TIME).Unix()
	tokenDetails.RefreshUuid = uuid.NewV4().String()

	refreshTokenClaims := jwt.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"exp":      tokenDetails.RtExpires,
	}

	var err error

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenDetails.RefreshToken, err = refresh.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return errors.New("can't create refresh token")
	}

	return nil
}

func CreateJWToken(user *entities.BaseUser) (*entities.JWTokenDetails, error) {
	JWTokenDetails := new(entities.JWTokenDetails)

	if err := CreateAccessToken(user, JWTokenDetails); err != nil {
		return nil, err
	}

	if err := CreateRefreshToken(user, JWTokenDetails); err != nil {
		return nil, err
	}

	return JWTokenDetails, nil
}
