package service

import (
	"errors"
	"fmt"

	"github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model

	Name string `json:"name" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
	Workout []Workout
}

func (exercise *Exercise) InsertExercise(db *gorm.DB, logger *utils.Logger) error {

	tx := db.Create(&exercise)
	logger.InfoLog.Printf("Inserting: %v  ", *exercise)
	if tx.Error != nil {
		logger.ErrorLog.Printf("Unexpected Error Occurred:%v \n", tx.Error)
		return errors.New("")

	}
	return nil

}

func GetAllExercise(db *gorm.DB, logger *utils.Logger) []Exercise {
	var e []Exercise
	db.Find(&e)
	logger.InfoLog.Print("Fetching All Exercises:")
	return e
}

func GetExerciseById(id string, db *gorm.DB, logger *utils.Logger) (*Exercise, *ServiceError) {
	var e Exercise
	result := db.First(&e, "id = ? ", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &ServiceError{
				Message:    fmt.Sprintf("There's No Exercise With ID:%s", id),
				StatusCode: 404,
			}
		} else {
			logger.ErrorLog.Printf("Something went wrong:%v", result.Error)
			return nil, &ServiceError{
				Message:    "Something went wrong try again",
				StatusCode: 500,
			}

		}

	}
	return &e, nil
}
