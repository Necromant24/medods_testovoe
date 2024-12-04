package interfaces

import "medods/auth-service/models"

type ITokensRepository interface {
	GetRefreshToken(userId string) (models.RefreshToken, error)
	SaveToken(token models.RefreshToken) error
}
