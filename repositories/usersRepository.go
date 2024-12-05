package repositories

import (
	"context"
	"fmt"
	"medods/auth-service/infrastructures"
	"medods/auth-service/models"

	"github.com/jackc/pgx/v5"
)

type UsersRepository struct {
	Conn *pgx.Conn
}

func (repo *UsersRepository) GetUserById(userId string) (models.User, error) {
	var err error

	repo.Conn, err = infrastructures.InitPgxConn()

	if err != nil {
		fmt.Println("Unable to connect to database: %v\n", err)
	}
	defer repo.Conn.Close(context.Background())

	user := models.User{}
	err = repo.Conn.QueryRow(context.Background(), "select id,name from users where id=$1", userId).Scan(&user.Id, &user.Name)
	if err != nil {
		fmt.Println("QueryRow failed: %v\n", err)
	}

	return user, err
}
