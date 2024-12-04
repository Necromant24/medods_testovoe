package interfaces

import "context"

type IDbHandler interface {
	Close(ctx context.Context) error
	QueryRow(ctx context.Context, sql string, args ...any) IRow
}

type IRow interface {
	Scan(dest ...any) error
}
