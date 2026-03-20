package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/common"
	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/repository"
)

type TodoService struct {
	repo repository.Repository
}

func NewTodoService(repo repository.Repository) *TodoService {
	return &TodoService{repo}
}

func (ts *TodoService) CreateTodo(ctx context.Context, todo Todo) error {

	slog.Info("SERVICE", "METHOD", "CreateTodo", "BODY", todo)

	if todo.Title == "" {
		return fmt.Errorf("Title is required!!")
	}

	err := ts.repo.CreateTodo(ctx, toTodoModel(todo))

	if err != nil {
		return err
	}

	return nil

}

func (ts *TodoService) GetTodos(ctx context.Context, filter common.Filter) ([]Todo, error) {

	slog.Info("SERVICE", "METHOD", "GetTodos")
	fmt.Println("")

	resp, err := ts.repo.GetTodos(ctx, filter)
	if err != nil {
		return nil, err
	}
	var todos []Todo
	for _, todo := range resp {
		todos = append(todos, FromTodoModel(todo))
	}

	return todos, nil
}

func (ts *TodoService) UpdateTodo(ctx context.Context, id int, todo Todo) error {

	slog.Info("SERVICE", "METHOD", "UpdateTodo", "ID", id, "BODY", todo)
	fmt.Println("")
	err := ts.repo.UpdateTodo(ctx, id, toTodoModel(todo))
	if err != nil {
		return err
	}
	return nil
}

func (ts *TodoService) DeleteTodo(ctx context.Context, id int) error {

	slog.Info("SERVICE", "METHOD", "DeleteTodo", "ID", id)
	fmt.Println("")

	if id <= 0 {
		return fmt.Errorf("Invalid Id!!")
	}
	err := ts.repo.DeleteTodo(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// dto

func toTodoModel(t Todo) repository.TodoModel {
	return repository.TodoModel{
		Title:       t.Title,
		Description: t.Description,
		IsComplete:  t.IsComplete,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func FromTodoModel(t repository.TodoModel) Todo {
	return Todo{
		Title:       t.Title,
		Description: t.Description,
		IsComplete:  t.IsComplete,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
