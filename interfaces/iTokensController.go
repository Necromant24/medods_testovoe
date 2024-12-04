package interfaces

import "net/http"

type ITokensController interface {
	GetTokensPair(res http.ResponseWriter, req *http.Request)
	RefreshTokensPair(res http.ResponseWriter, req *http.Request)
}
