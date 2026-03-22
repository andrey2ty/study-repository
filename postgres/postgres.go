package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnetctionToPostgres(ctx context.Context) (*pgx.Conn, error) {
	conn := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, conn)
}
