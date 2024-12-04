package main

import (
	"context"
	"fmt"
	"medods/auth-service/config"
	"medods/auth-service/controllers"
	"medods/auth-service/interfaces"
	"medods/auth-service/repositories"
	"medods/auth-service/services"

	"github.com/jackc/pgx/v5"
)

type IServiceContainer interface {
	InjectUsersController() interfaces.IUsersController
}

func InjectUsersController() interfaces.IUsersController {

	conn, err := pgx.Connect(context.Background(), config.GetConfig().DbConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	usersRepository := &repositories.UsersRepository{conn}
	usersService := &services.UsersService{usersRepository}
	usersController := &controllers.UsersController{usersService}

	return usersController
}

func InjectTokensController() interfaces.ITokensController {
	conn, err := pgx.Connect(context.Background(), config.GetConfig().DbConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	usersRepository := &repositories.UsersRepository{conn}
	usersService := &services.UsersService{usersRepository}
	tokensRepo := &repositories.TokensRepository{conn}
	tokensService := &services.TokensService{tokensRepo}
	tokensController := &controllers.TokensController{usersService, tokensService}

	return tokensController
}
