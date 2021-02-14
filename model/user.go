package model

type User struct {
	ID       int
	Email    string `json:"email" validate:"required,email,unique"` //Check unique field (unique=table:field)
	Password string `json:"password" validate:"gte=5,lte=50"`
	TokenJWT string
}
