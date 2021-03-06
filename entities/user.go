package entities

import (
	"gorm.io/gorm"
)

type BaseUser struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserReqEntry struct {
	ID       uint
	Username string
	Password string
	Email    string
}
