package mapper

import (
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/models"
)

func MapUserDtoToUser(user dto.User) models.User {
	return models.User{
		Name:     user.Name,
		Email:    user.Email,
		Username: &user.Username,

		Birthday: user.Birthday,
		Password: user.Password,
	}
}

func MapUserToUserDto(user models.User) dto.User {
	return dto.User{
		Name:     user.Name,
		Email:    user.Email,
		Birthday: user.Birthday,
		Password: user.Password,
		Username: *user.Username,
	}
}
