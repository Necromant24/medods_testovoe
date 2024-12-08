package controllers

import (
	"encoding/json"
	"fmt"
	"medods/auth-service/interfaces"
	"medods/auth-service/models"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
)

type TokensController struct {
	interfaces.IUsersService
	interfaces.ITokensService
}

func (controller *TokensController) RefreshTokensPair(w http.ResponseWriter, r *http.Request) {

	var token models.RefreshToken
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("error in getting ip address")
	}
	token.UserIp = ip

	access, refresh, err := controller.ITokensService.RefreshTokens(token.AccessToken, token.Token)

	ren := render.New()

	var dict map[string]string = make(map[string]string)
	dict["accessToken"] = access
	dict["refreshToken"] = refresh

	status := http.StatusOK

	if access == "" {
		status = http.StatusBadRequest
	}

	ren.JSON(w, status, dict)

}

func (controller *TokensController) GetTokensPair(res http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "userId")

	r := render.New()

	user, err := controller.IUsersService.GetUserById(userId)
	if err != nil {
		fmt.Println(err)
	}

	if user.Id == "" {
		r.JSON(res, http.StatusNotFound, userId)
		return
	}

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		fmt.Println("error in getting ip address")
	}

	access, refresh, err := controller.ITokensService.GetTokensPair(user.Id, ip)
	if err != nil {
		fmt.Println("error in getting ip address")
	}

	var dict map[string]string = make(map[string]string)
	dict["accessToken"] = access
	dict["refreshToken"] = refresh

	r.JSON(res, http.StatusOK, dict)
}
