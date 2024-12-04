package interfaces

import "medods/auth-service/models"

type IUsersRepository interface {
	GetUserById(userId string) (models.User, error)
}
