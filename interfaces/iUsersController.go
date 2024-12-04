package interfaces

import "net/http"

type IUsersController interface {
	GetUserById(res http.ResponseWriter, req *http.Request)
}
