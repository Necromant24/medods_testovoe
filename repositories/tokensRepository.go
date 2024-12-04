package repositories

import (
	"context"
	"fmt"
	"medods/auth-service/models"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TokensRepository struct {
	Conn *pgx.Conn
}

// gets bcrypt hash, because we save in bcrypt hash
func (repo *TokensRepository) GetRefreshToken(userId string) (models.RefreshToken, error) {
	var err error
	token := models.RefreshToken{}
	err = repo.Conn.QueryRow(context.Background(), "select id,token, userId, userIp from tokens where userId=$1",
		userId).Scan(&token.Id, &token.Token, &token.UserId, &token.UserIp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(token)
	return token, err
}

func (repo *TokensRepository) SaveToken(token models.RefreshToken) error {

	uuid := uuid.New()
	sql := `
			INSERT INTO tokens (id, token, userId, userIp)
			VALUES ($1,$2,$3,$4)
			ON CONFLICT (userId) DO UPDATE
			SET 
			token = EXCLUDED.token,
			userIp = EXCLUDED.userIp;
	`

	_, err := repo.Conn.Exec(context.Background(), sql, uuid, token.Token, token.UserId, token.UserIp)
	return err
}
