package validators

import (
	"fiberJWTAuth/entities"
	"fiberJWTAuth/search"
	"net/mail"
	"regexp"
)

var (
	usernameRegex, _ = regexp.Compile("^[a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[a-zA-Z0-9]$")
	passwordRegex, _ = regexp.Compile("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$")
)

func VerifyUsername(user *entities.UserReqEntry) bool {
	return user.Username != "" && usernameRegex.MatchString(user.Username)
}

func VerifyPassword(user *entities.UserReqEntry) bool {
	return user.Username != "" && passwordRegex.MatchString(user.Password)
}

func VerifyEmail(user *entities.UserReqEntry) bool {
	_, err := mail.ParseAddress(user.Email)
	return err == nil
}

func IsExist(user *entities.UserReqEntry) bool {
	_, err := search.GetUser(user)
	if err != nil {
		return false
	}
	return true
}

func VerifyReqUser(reqUser *entities.UserReqEntry) bool {
	return VerifyUsername(reqUser) &&
		VerifyPassword(reqUser) &&
		VerifyEmail(reqUser) &&
		(!IsExist(reqUser))
}

func IsCorrectLoginData(reqUser *entities.UserReqEntry, user *entities.BaseUser) bool {
	return reqUser.Username == user.Username && reqUser.Password == user.Password
}
