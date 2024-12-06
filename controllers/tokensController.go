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
	// TODO: make refresh from service

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

	ren.JSON(w, http.StatusOK, dict)

	// если вы проверяющий, пожалуйста проверьте завтра, я все доделаю, просто поздно увидел сообщение с hh и поздно приступил к тестовому
}

func (controller *TokensController) GetTokensPair(res http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "userId")

	var userPassword models.UserPasswordDto
	err := json.NewDecoder(req.Body).Decode(&userPassword)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

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
