package mapper

import (
	"github.com/zYasser/MyFitness/internal/user/dto"
	"github.com/zYasser/MyFitness/internal/user/service"
)

func MapUserDtoToUser(user dto.User) service.User {
	return service.User{
		Name:     user.Name,
		Email:    &user.Email,
		Username: &user.Username,

		Birthday: user.Birthday,
		Password: user.Password,
	}
}

func MapUserToUserDto(user service.User) dto.User {
	return dto.User{
		Name:     user.Name,
		Email:    *user.Email,
		Birthday: user.Birthday,
		Password: user.Password,
		Username: *user.Username,
	}
}
