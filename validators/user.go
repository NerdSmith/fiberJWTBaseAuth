package validators

import (
	"fiberJWTAuth/entities"
	"net/mail"
	"regexp"
)

var (
	usernameRegex, _ = regexp.Compile("^[a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[a-zA-Z0-9]$")
	passwordRegex, _ = regexp.Compile("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$")
)

func VerifyUsername(user *entities.UserReqEntry) bool {
	return usernameRegex.MatchString(user.Username)
}

func VerifyPassword(user *entities.UserReqEntry) bool {
	return passwordRegex.MatchString(user.Password)
}

func VerifyEmail(user *entities.UserReqEntry) bool {
	_, err := mail.ParseAddress(user.Email)
	return err == nil
}
