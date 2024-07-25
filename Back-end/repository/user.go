package repository

import (
	"fmt"
	"time"

	"github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
)
var logger = utils.GetLogger()

type User struct {
	gorm.Model

    ID       uint   `gorm:"primaryKey;autoIncrement"`
    Name     string
    Email    string
    Birthday *time.Time
    Password string
}

func (user *User) CreateUser(db *gorm.DB) error {
	tx:=db.Create(&user)
	err:=tx.Error
	fmt.Println(err)
	tx.Commit()
	if(err!=nil){
		logger.ErrorLog.Printf("Error While Inserting %v, Error: %v\n", user, err)
	}
	return nil
}
