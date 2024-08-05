package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET_KEY = GetEnv("JWT_SECRET_KEY")
var JWT_TTL = GetEnv("JWT_TTL")



func CreateToken(username string ) (string, error) {

        secretKey :=[]byte(JWT_SECRET_KEY)
        JWT_TTL, err := strconv.Atoi(JWT_TTL)
        if(err!=nil){
            log.Println("JWT_TTL environment variable should be a number ")
            JWT_TTL=10
        }
	    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Second * time.Duration(JWT_TTL)).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    
    if err != nil {
        log.Println("The creation of JWT has failed Error %s \n" , err)
    return "", err
    }

 return tokenString, nil
}


func VerifyToken(tokenString string) error {
    secretKey :=[]byte(JWT_SECRET_KEY)

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
       return secretKey, nil
    })
   
    if err != nil {
       return err
    }
   
    if !token.Valid {
       return fmt.Errorf("invalid token")
    }
   
    return nil
 }
 