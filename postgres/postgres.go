package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnetctionToPostgres(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:0710@localhost:5432/postgres")
}
