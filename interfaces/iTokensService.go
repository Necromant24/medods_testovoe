package interfaces

type ITokensService interface {
	GetTokensPair(userId string, userIp string) (string, string, error)
}
