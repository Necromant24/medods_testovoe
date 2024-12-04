package services

import (
	"medods/auth-service/interfaces"
	"medods/auth-service/models"
)

type UsersService struct {
	UsersRepository interfaces.IUsersRepository
}

func (service *UsersService) GetUserById(userId string) (models.User, error) {

	return service.UsersRepository.GetUserById(userId)
}
