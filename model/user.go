package model

import "websocket/command"

type User struct {
	Email    string
	password string
	tokenJWT string
	command  command.Command
}

func Authentication(Email string, Password string) *User {
	return &User{
		Email:    "13123",
		password: "123123",
		tokenJWT: "weeqewq",
	}
}
