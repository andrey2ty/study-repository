package dockerpractice

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PeopleModel struct {
	Id       int
	FullName string
	Position string
}

func NewPeopleModel(id int, fullname string, position string) PeopleModel {
	return PeopleModel{
		Id:       id,
		FullName: fullname,
		Position: position,
	}
}

func DeletePeople(conn *pgx.Conn, ctx context.Context, id int) error {
	sql := `
	DELETE FROM people
	WHERE id = $1
	`
	_, err := conn.Exec(ctx, sql, id)

	return err
}

func InsertPeople(conn *pgx.Conn, ctx context.Context, people PeopleModel) error {
	sql := `
	INSERT INTO people (full_name, position)
	VALUES ($1, $2);
	`

	_, err := conn.Exec(ctx, sql, people.FullName, people.Position)

	return err
}

func SelectPeople(conn *pgx.Conn, ctx context.Context) ([]PeopleModel, error) {
	sql := `
		SELECT * FROM people;
	`
	rows, err := conn.Query(ctx, sql)

	if err != nil {
		return []PeopleModel{}, nil
	}

	arr := make([]PeopleModel, 0)

	for rows.Next() {
		var people PeopleModel

		if err := rows.Scan(&people.Id, &people.FullName, &people.Position); err != nil {
			return []PeopleModel{}, nil
		}

		arr = append(arr, people)
	}

	return arr, nil
}
