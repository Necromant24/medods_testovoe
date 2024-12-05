package models

type UserDto struct {
	Id           string `json:"id"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserPasswordDto struct {
	UserPassword string `json:"userPassword"`
}
