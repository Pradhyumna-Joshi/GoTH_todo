package repository

import (
	"context"

	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/common"
)

type Repository interface {
	CreateTodo(context.Context, TodoModel) (TodoModel, error)
	GetTodos(context.Context, common.Filter) ([]TodoModel, error)
	ToggleTodo(context.Context, int) (TodoModel, error)
	UpdateTodo(context.Context, int, TodoModel) (TodoModel, error)
	DeleteTodo(context.Context, int) error
}
