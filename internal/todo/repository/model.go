package repository

import "time"

type TodoModel struct {
	Id          int32
	Title       string
	Description string
	IsComplete  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
