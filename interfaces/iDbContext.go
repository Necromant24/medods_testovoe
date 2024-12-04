package interfaces

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type IDbContext interface {
	Connect(ctx context.Context, connString string) (*pgx.Conn, error)
}
