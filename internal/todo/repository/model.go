package repository

import "time"

type TodoModel struct {
	Id          int
	Title       string
	Description string
	IsComplete  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
