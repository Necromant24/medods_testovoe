package infrastructures

import (
	"context"
	"fmt"
	"medods/auth-service/config"

	"github.com/jackc/pgx/v5"
)

func InitPgxConn() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), config.GetConfig().DbConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	return conn, err
}
