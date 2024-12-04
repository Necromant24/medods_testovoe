package interfaces

import "medods/auth-service/models"

type IUsersService interface {
	GetUserById(userId string) (models.User, error)
}
