package mapper

import (
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/models"
)

func MapToWorkout(dto dto.WorkoutCreateDTO) *models.Workout {
	return &models.Workout{
		Name:          dto.Name,
		Day:           dto.Day,
		RepLowerBound: dto.RepLowerBound,
		RepUpperBound: dto.RepUpperBound,
		Description:   dto.Description,
		ExerciseID:    dto.ExerciseID,
		ProgramId:     dto.ProgramId,
	}
}
