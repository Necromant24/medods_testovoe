package controllers

import (
	"fmt"
	"medods/auth-service/interfaces"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
)

type TokensController struct {
	interfaces.IUsersService
	interfaces.ITokensService
}

func (controller *TokensController) RefreshTokensPair(res http.ResponseWriter, req *http.Request) {
	// TODO: make refresh from service
}

func (controller *TokensController) GetTokensPair(res http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "userId")

	r := render.New()

	user, err := controller.IUsersService.GetUserById(userId)
	if err != nil {
		fmt.Println(err)
	}

	r.JSON(res, http.StatusOK, user)
}
