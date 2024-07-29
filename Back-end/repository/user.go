package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
)
var logger = utils.GetLogger()

type User struct {
	gorm.Model

    ID       uint   `gorm:"primaryKey;autoIncrement"`
    Name     string
    Email    *string `gorm:"unique"`
	Username    *string `gorm:"unique"`
    Birthday *time.Time
    Password string
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