package main

import (
	"context"
	"fmt"
	"study/postgres"
	feature_sql "study/postgres/featureSql"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := postgres.ConnetctionToPostgres(ctx)
	if err != nil {
		panic(err)
	}
	if err := feature_sql.CreateTable(conn, ctx); err != nil {
		panic(err)
	}

	tasks, err := feature_sql.SelectRow(conn, ctx)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.Id == 8 {
			task.Completed = true
			now := time.Now()
			task.Completed_at = &now

			if err := feature_sql.UpdateRow(conn, ctx, task); err != nil {
				panic(err)
			}

			break
		}
	}

	fmt.Println("succeed!")
}
