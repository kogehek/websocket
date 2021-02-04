package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
}

func ExampleNewWithClaims_standardClaims() string {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Sub  string `json:"sub"`
		Name string `json:"name"`
		Iat  string `json:"iat"`
		jwt.StandardClaims
	}
	claims := MyCustomClaims{
		"bar",
		"John Doe",
		"1516239022",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Println(ss)
	return ss
}

func Time(tokenString string) {
	type MyCustomClaims struct {
		Sub  string `json:"sub"`
		Name string `json:"name"`
		Iat  string `json:"iat"`
		jwt.StandardClaims
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		t := time.Unix(claims.StandardClaims.ExpiresAt, 0)
		fmt.Println(t.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(err)
	}
}
