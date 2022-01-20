package main

type User struct {
	UID      uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var TUser = User{
	UID:      1,
	Username: "NS",
	Password: "1029",
}
