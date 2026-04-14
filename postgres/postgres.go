package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnetctionToPostgres(ctx context.Context) (*pgx.Conn, error) {
	godotenv.Load()

	conn := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, conn)
}
