package service

import (
	"errors"

	"github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model

	Name string `json:"name" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
}



func (exercise *Exercise) InsertExercise(db *gorm.DB , logger *utils.Logger) error{

	tx:=db.Create(&exercise)
	logger.InfoLog.Printf("Inserting: %v  " , *exercise )
	if(tx.Error != nil){
		logger.ErrorLog.Printf("Unexpected Error Occurred:%v \n" , tx.Error)
		return errors.New("")

	}
	return nil

}



func GetAllExercise(db *gorm.DB , logger *utils.Logger) []Exercise{
	var e []Exercise
	db.Find(&e)
	logger.InfoLog.Print("Fetching All Exercises:")
	return e
}