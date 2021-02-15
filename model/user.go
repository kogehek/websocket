package model

type User struct {
	ID       int
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"gte=5,lte=50"`
}
