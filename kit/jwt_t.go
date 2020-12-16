package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClient struct {
	jwt.StandardClaims
	Name string `json:"name"`
}

func main() {
	sce := []byte("123abc")

	user := &UserClient{Name: "abc"}
	user.ExpiresAt = time.Now().Add(time.Second * 5).Unix() //token 有效期 5秒
	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	token, _ := tokenObject.SignedString(sce)
	fmt.Println(token)

	getToken, _ := jwt.ParseWithClaims(token, user, func(token *jwt.Token) (interface{}, error) {
		return sce, nil
	})

	if getToken.Valid {
		fmt.Println(getToken.Claims.(*UserClient).Name)
	}
}
