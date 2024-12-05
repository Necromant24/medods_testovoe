package interfaces

type ITokensService interface {
	GetTokensPair(userId string, userIp string) (string, string, error)
	RefreshTokens(accessToken string, refreshToken string) (string, string, error)
}
