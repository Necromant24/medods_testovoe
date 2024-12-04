package controllers

import (
	"fmt"
	"medods/auth-service/interfaces"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
)

type UsersController struct {
	interfaces.IUsersService
}

func (usersController *UsersController) GetUserById(res http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "userId")

	r := render.New()

	user, err := usersController.IUsersService.GetUserById(userId)
	if err != nil {
		fmt.Println(err)
	}

	r.JSON(res, http.StatusOK, user)

}
