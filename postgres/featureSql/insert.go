package feature_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(conn *pgx.Conn, ctx context.Context, task TaskModel) error {
	sql := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1, $2, $3, $4);
	`
	_, err := conn.Exec(ctx, sql,
		task.Title,
		task.Description,
		task.Completed,
		task.Created_at,
	)

	return err
}
