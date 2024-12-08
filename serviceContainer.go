package main

import (
	"medods/auth-service/controllers"
	"medods/auth-service/interfaces"
	"medods/auth-service/repositories"
	"medods/auth-service/services"
)

type IServiceContainer interface {
	InjectUsersController() interfaces.IUsersController
}

func InjectUsersController() interfaces.IUsersController {

	usersRepository := &repositories.UsersRepository{}
	usersService := &services.UsersService{usersRepository}
	usersController := &controllers.UsersController{usersService}

	return usersController
}

func InjectTokensController() interfaces.ITokensController {

	usersRepository := &repositories.UsersRepository{}
	usersService := &services.UsersService{usersRepository}
	tokensRepo := &repositories.TokensRepository{}
	tokensService := &services.TokensService{tokensRepo}
	tokensController := &controllers.TokensController{usersService, tokensService}

	return tokensController
}
