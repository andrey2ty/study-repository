package feature_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(conn *pgx.Conn, ctx context.Context, sliceId []int) error {
	sql := `
	DELETE FROM tasks
	WHERE id = ANY($1)
	`
	_, err := conn.Exec(ctx, sql, sliceId)

	return err
}
