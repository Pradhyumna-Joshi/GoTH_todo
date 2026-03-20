package service

import (
	"context"

	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/common"
)

type Service interface {
	CreateTodo(context.Context, Todo) error
	GetTodos(context.Context, common.Filter) ([]Todo, error)
	UpdateTodo(context.Context, int, Todo) error
	DeleteTodo(context.Context, int) error
}
