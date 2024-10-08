package mapper

import (
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/service"
)

func DtoToExercise(dto dto.Exercise) *service.Exercise {
	return &service.Exercise{
		Name: dto.Name,
		Type: dto.Type,
	}
}