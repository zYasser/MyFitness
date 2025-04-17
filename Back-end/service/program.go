package service

import (
	"errors"
	"fmt"

	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model

	ID          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Description *string
	Workout     []Workout
}

func dtoToProgram(dto dto.Program) *Program {
	return &Program{
		Name:        dto.Name,
		Description: &dto.Name,
	}

}

func InsertProgram(db *gorm.DB, logger *utils.Logger, request dto.Program) error {
	logger.InfoLog.Println("Saving Program Into Database")
	program := dtoToProgram(request)
	tx := db.Create(&program)
	if tx.Error != nil {
		logger.ErrorLog.Printf("Unexpected Error Occurred:%v \n", tx.Error)
		return errors.New("")

	}
	return nil

}

func AddWorkoutToProgram(db *gorm.DB, logger *utils.Logger, workouts []Workout, id uint) error {
	for _, element := range workouts {
		element.ProgramId = id
	}
	tx := db.Create(&workouts)
	if tx.Error != nil {
		logger.ErrorLog.Printf("Unexpected Error Occurred:%v \n", tx.Error)
		return errors.New("")

	}
	return nil

}

func GetAllProgram(db *gorm.DB, logger *utils.Logger) []Program {
	var result []Program
	tx := db.Find(&result)
	fmt.Println(tx.Error)
	return result
}

func GetProgramById(db *gorm.DB, logger *utils.Logger, id int64) (*Program, *ServiceError) {
	var result Program
	tx := db.First(&result).Where("id = ? ", id)
	if tx.Error != nil {
		return nil, &ServiceError{
			Message:    tx.Error.Error(),
			StatusCode: 404,
		}
	}

	return &result, nil
}
