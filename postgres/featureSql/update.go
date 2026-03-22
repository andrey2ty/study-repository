package feature_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(conn *pgx.Conn, ctx context.Context, task TaskModel) error {
	sql := `
	UPDATE tasks
	SET title = $1, description = $2 , completed = $3, created_at = $4, completed_at = $5
	WHERE id = $6;
	`

	_, err := conn.Exec(
		ctx,
		sql,
		task.Title,
		task.Description,
		task.Completed,
		task.Created_at,
		task.Completed_at,
		task.Id,
	)

	return err
}
