package model

import (
	"websocket/valid"
)

type User struct {
	ID       int
	Email    string `validate:"required,gvna"`
	Password string `validate:"gte=5,lte=50"`
	TokenJWT string
}

func NewUser(Email string, encrypted_password string) (*User, error) {
	user := &User{
		Email:    Email,
		Password: encrypted_password,
	}
	err := valid.Model(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
