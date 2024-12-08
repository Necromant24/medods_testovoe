package models

type RefreshToken struct {
	Id          string `json:"id"`
	Token       string `json:"refreshToken"`
	UserId      string `json:"userId"`
	UserIp      string `json:"userIp"`
	AccessToken string `json:"accessToken"`
}
