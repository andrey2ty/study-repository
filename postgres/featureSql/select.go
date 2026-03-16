package feature_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRow(conn *pgx.Conn, ctx context.Context) ([]TaskModel, error) {
	sql := `
	SELECT * FROM tasks;
	`

	rows, err := conn.Query(ctx, sql)

	if err != nil {
		return []TaskModel{}, nil
	}
	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		if err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Created_at,
			&task.Completed_at,
		); err != nil {
			return []TaskModel{}, nil
		}

		tasks = append(tasks, task)

	}

	return tasks, nil
}

func printTask(task TaskModel) {

	fmt.Println("id:", task.Id)
	fmt.Println("title:", task.Title)
	fmt.Println("description:", task.Description)
	fmt.Println("completed:", task.Completed)
	fmt.Println("created_at:", task.Created_at)
	fmt.Println("completed_at:", task.Completed_at)
	fmt.Println("------------------")
}
