package handler

import (
	jwt "github.com/dgrijalva/jwt-go"
	"log"
)

func ValidateAccessToken(accessToken string, privateKey string) (bool, error) {
	var flag bool
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err == nil && token.Valid {
		log.Println("Access Token is valid")
		flag = true
		return flag, nil
	} else {
		log.Println("Access Token is invalid")
		flag = false
		return flag, err
	}
}
