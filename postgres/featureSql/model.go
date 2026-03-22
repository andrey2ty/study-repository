package feature_sql

import "time"

type TaskModel struct {
	Id           int
	Title        string
	Description  string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
}
