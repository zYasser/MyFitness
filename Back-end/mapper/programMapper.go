package mapper

import (
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/models"
)

func DtoToProgram(dto dto.Program) *models.Program {
	return &models.Program{
		Name:        dto.Name,
		Description: &dto.Name,
	}

}
