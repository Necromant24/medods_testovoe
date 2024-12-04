package infrastructures

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PgxHandler struct {
	Conn *pgx.Conn
}

type PgxRow struct {
}

func (dbHandler *PgxHandler) Close(ctx context.Context) error {
	return dbHandler.Close(context.Background())
}

func (dbHandler *PgxHandler) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return dbHandler.Conn.QueryRow(ctx, sql, args)
}
