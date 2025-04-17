package service

import (
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/mapper"
	"github.com/zYasser/MyFitness/models"
	"github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
)

func createWorkout(db *gorm.DB, logger *utils.Logger, dto dto.WorkoutCreateDTO) (*models.Workout, *ServiceError) {
	workout := mapper.MapToWorkout(dto)

	if err := db.Create(workout).Error; err != nil {
		return nil, &ServiceError{
			Message:    "Failed to create workout",
			StatusCode: 500,
		}
	}

	return workout, nil
}
