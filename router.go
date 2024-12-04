package main

// import (
// 	"sync"

// 	"github.com/go-chi/chi/v5"
// )

// type IChiRouter interface {
// 	InitRouter() *chi.Mux
// }

// type router struct{}

// func (router *router) InitRouter() *chi.Mux {

// 	usersController := ServiceContainer().InjectUsersController()

// 	r := chi.NewRouter()
// 	r.HandleFunc("/users/{userId}", usersController.GetUserById)

// 	return r
// }

// var (
// 	m          *router
// 	routerOnce sync.Once
// )

// func ChiRouter() IChiRouter {
// 	if m == nil {
// 		routerOnce.Do(func() {
// 			m = &router{}
// 		})
// 	}
// 	return m
// }
