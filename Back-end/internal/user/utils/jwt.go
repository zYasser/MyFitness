package utils

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var JWT_SECRET_KEY = GetEnv("JWT_SECRET_KEY")
var JWT_TTL = GetEnv("JWT_TTL")

type JwtExpireTokenErr struct {
}

func (err JwtExpireTokenErr) Error() string {
	return "Token is Expired"
}

func CreateToken(username string) (string, string, error) {

	secretKey := []byte(JWT_SECRET_KEY)
	JWT_TTL, err := strconv.Atoi(JWT_TTL)
	refresh := uuid.New().String()
	if err != nil {
		log.Println("JWT_TTL environment variable should be a number ")
		JWT_TTL = 10
	}
    _=JWT_TTL
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour*time.Duration(JWT_TTL)).Unix(),
			"refresh":  refresh,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("The creation of JWT has failed Error %s \n", err)
		return "", "", err
	}

	return tokenString, refresh, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	secretKey := []byte(JWT_SECRET_KEY)
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil && errors.Is(err, jwt.ErrTokenExpired) {
        fmt.Println("T")
		return claims , JwtExpireTokenErr{}

	}
	if err != nil {
		return nil, err
	}
	return claims, err
}
