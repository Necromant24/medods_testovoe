package main

// import (
// 	"context"
// 	"fmt"
// 	"medods/auth-service/config"
// 	"medods/auth-service/repositories"
// 	"net/http"
// 	"time"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// 	"github.com/jackc/pgx/v5"
// 	httpSwagger "github.com/swaggo/http-swagger/v2"

// 	_ "medods/auth-service/docs"
// )

// // @title       My API
// // @version     1.0
// // @description This is a sample API.
// // @host        localhost:3001
// // @BasePath    /api/v1
// func main() {

// 	config.LoadConfiguration()

// 	port := config.GetConfig().HostPort
// 	// dbConn := config.GetConfig().DbConnectionString

// 	// urlExample := "postgres://username:password@localhost:5432/database_name"
// 	var err error
// 	var usersRepository repositories.UsersRepository
// 	usersRepository.Conn, err = pgx.Connect(context.Background(), config.GetConfig().DbConnectionString)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	r := chi.NewRouter()

// 	r.Use(middleware.Logger)
// 	// processing should be stopped.
// 	r.Use(middleware.Timeout(60 * time.Second))

// 	r.Get("/swagger/*", httpSwagger.Handler(
// 		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port)), //The url pointing to API definition
// 	))

// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("welcome"))
// 	})

// 	r.Get("/greet", Greet)

// 	r.Get("/user/{userId}")

// 	http.ListenAndServe(fmt.Sprintf(":%d", port), r)

// }

// // @Summary Get Access and Refresh tokens by userId
// // @Tags users
// // @Accept json
// // @Produce json
// // @Param id path string true "User ID"
// // @Router /access/{id} [get]
// func GetAccess(userId string) User {
// 	return User{}
// }

// // @Summary Refresh tokensR
// // @Tags users
// // @Accept json
// // @Produce json
// // @Router /refresh [post]
// func Refresh(user User) User {
// 	return User{}
// }

// // Greet example
// // @Summary Get user by ID
// // @Description Get details of a user by ID
// // @Tags users
// // @Accept  json
// // @Produce  json
// // @Param   id   path      int     true  "User ID"
// // @Success 200  {object}  User
// // @Failure 404  {object}  Error
// // @Router  /users/{id} [get]
// func Greet(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("welcome greet"))
// }

// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// type Error struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }
