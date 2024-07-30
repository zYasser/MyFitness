package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
)
var logger = utils.GetLogger()

type User struct {
	gorm.Model

    ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    Name     string `json:"name"`
    Email    *string `gorm:"unique" json:"email"`
	Username  *string `gorm:"unique" json:"username"`
    Birthday *time.Time `json:"birthday"`
    Password string `json:"-"`
}

func (user *User) CreateUser(db *gorm.DB) error {
    fmt.Println(user.Email)
	hashedPassword, err :=utils.HashPassword(user.Password)
	if(err !=nil){
		logger.ErrorLog.Printf("Failed To hash the password")
		return errors.New("")

	}
	user.Password=hashedPassword
	tx:=db.Create(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		if strings.Contains(tx.Error.Error(), "duplicate key value violates unique constraint") {
			col := utils.ExtractColumn(tx.Error.Error())
			return fmt.Errorf("%s Already Exist" , col)
		} else {
			logger.ErrorLog.Printf("Unexpected Error Occurred:%v \n" , tx.Error)
			return errors.New("")
		}
	}
    return nil
}

func ValidateUser(db *gorm.DB ,  credential dto.UserLogin ) error{
	logger.InfoLog.Printf("User:%s trying to login  \n " , credential.Username)

	var user User
	result:=db.Where("username= ?" , credential.Username).First(&user)
	if(result.Error !=nil){	
		if(errors.Is(result.Error, gorm.ErrRecordNotFound)){
			logger.ErrorLog.Printf("%s User doesn't exist" ,credential.Username)
			return fmt.Errorf("%s User doesn't exist" ,credential.Username )
		}
		
		logger.ErrorLog.Printf("Unexpected Error Occurred:%v \n" , result.Error)
		return errors.New("")

		
	}
	if(!utils.CheckPasswordHash(credential.Password, user.Password)){
		logger.ErrorLog.Printf("User:%s tried to login with wrong password \n" ,credential.Username )
		return errors.New("check you username and password")

	}
	logger.InfoLog.Printf("User:%s successfully logged in \n " , credential.Username)
	return nil 
}

