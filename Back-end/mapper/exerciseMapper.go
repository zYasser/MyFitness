package mapper

import (
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/models"
)

func DtoToExercise(dto dto.Exercise) *models.Exercise {
	return &models.Exercise{
		Name: dto.Name,
		Type: dto.Type,
	}
}