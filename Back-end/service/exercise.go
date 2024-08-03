package service

import (
	"errors"

	"github.com/zYasser/MyFitness/middleware"
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
}



func (exercise *Exercise) InsertExercise(db *gorm.DB , logger *middleware.Logger) error{

	tx:=db.Create(&exercise)
	logger.InfoLog.Printf("Inserting: %v \n " , *exercise )
	if(tx.Error != nil){
		logger.ErrorLog.Printf("Unexpected Error Occurred:%v \n" , tx.Error)
		return errors.New("")

	}
	return nil

}