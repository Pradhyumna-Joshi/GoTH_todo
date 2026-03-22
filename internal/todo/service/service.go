package service

import (
	"context"

	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/common"
)

type Service interface {
	CreateTodo(context.Context, common.Todo) (common.Todo, error)
	GetTodos(context.Context, common.Filter) ([]common.Todo, error)
	ToggleTodo(context.Context, int) (common.Todo, error)
	UpdateTodo(context.Context, int, common.Todo) (common.Todo, error)
	DeleteTodo(context.Context, int) error
}
