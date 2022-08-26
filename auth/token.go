package auth

import (
	"Kirtasite/services"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(userId string) string {
	fmt.Println("Start")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["user"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	tokenString, _ := token.SignedString([]byte("I-Am-not-user-this-key"))
	return tokenString
}

func IsValid(data string) (bool, map[string]interface{}) {
	tkn, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		return []byte("I-Am-not-user-this-key"), nil
	})

	if err != nil {
		return false, services.SendMessage(services.NotLogin)
	}

	if !tkn.Valid {
		return false, services.SendMessage(services.Unauthorized)
	} else {
		return true, nil
	}
}

func TokenParser(data string) string {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(data, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("I-Am-not-user-this-key"), nil
	})
	userId := fmt.Sprintf("%v", claims["user"])
	return userId
}
