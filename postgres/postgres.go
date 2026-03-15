package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConnetctionToPostgres() {
	ctx := context.Background()
	connect, err := pgx.Connect(ctx, "postgres://postgres:0710@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := connect.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Успешное подключение")
}
