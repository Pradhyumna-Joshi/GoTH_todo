package service

import "time"

type Todo struct {
	Id          int
	Title       string
	Description string
	IsComplete  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
