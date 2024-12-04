package main

import (
	"medods/auth-service/config"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {

	config.LoadConfiguration()

	port := config.GetConfig().HostPort

	usersController := InjectUsersController()
	tokensController := InjectTokensController()

	r := chi.NewRouter()
	r.HandleFunc("/users/{userId}", usersController.GetUserById)
	r.HandleFunc("/getpair/{userId}", tokensController.GetTokensPair)
	r.HandleFunc("/refresh", tokensController.RefreshTokensPair)

	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
