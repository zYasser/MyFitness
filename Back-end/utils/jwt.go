package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zYasser/MyFitness/middleware"
)

var JWT_SECRET_KEY = GetEnv("JWT_SECRET_KEY")
var JWT_TTL = GetEnv("JWT_TTL")



func CreateToken(username string , log *middleware.Logger ) (string, error) {

        secretKey :=[]byte(JWT_SECRET_KEY)
        JWT_TTL, err := strconv.Atoi(JWT_TTL)
        if(err!=nil){
            log.ErrorLog.Println("JWT_TTL environment variable should be a number ")
            JWT_TTL=10
        }
	    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * time.Duration(JWT_TTL)).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    
    if err != nil {
        log.ErrorLog.Printf("The creation of JWT has failed Error %s \n" , err)
    return "", err
    }

 return tokenString, nil
}


