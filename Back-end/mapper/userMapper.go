package mapper

import (
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/repository"
)

func MapParametersToUser(user dto.User) repository.User {
	return repository.User{
		Name:     user.Name,
		Email:    &user.Email,
		Username:    &user.Username,

		Birthday: user.Birthday,
		Password: user.Password,
	}
}

func MapUserToParameters(user repository.User) dto.User {
	return dto.User{
		Name:     user.Name,
		Email:    *user.Email,
		Birthday: user.Birthday,
		Password: user.Password,
		Username: *user.Username,

	}
}