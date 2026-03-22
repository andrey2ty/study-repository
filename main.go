package main

import (
	"context"
	"fmt"
	"study/postgres"
	feature_sql "study/postgres/featureSql"
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

	fmt.Println("succeed!")
}
