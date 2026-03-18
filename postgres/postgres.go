package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnetctionToPostgres(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
}
