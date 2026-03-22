package repository

import (
	"context"
	"strings"
	"time"

	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostGresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool}
}

func (pgs *PostgresRepository) CreateTodo(ctx context.Context, todo TodoModel) (TodoModel, error) {

	var newTodo TodoModel
	err := pgs.pool.QueryRow(ctx, "INSERT INTO todos(title,description,is_complete,created_at,updated_at) VALUES($1,$2,$3,$4,$5) RETURNING *", todo.Title, todo.Description, todo.IsComplete, todo.CreatedAt, todo.UpdatedAt).Scan(&newTodo.Id, &newTodo.Title, &newTodo.Description, &newTodo.IsComplete, &newTodo.CreatedAt, &newTodo.UpdatedAt)
	if err != nil {
		return TodoModel{}, err
	}

	return newTodo, nil
}

func (pgs *PostgresRepository) GetTodos(ctx context.Context, filter common.Filter) ([]TodoModel, error) {

	query := "SELECT * FROM todos"

	if strings.ToLower(filter.Completed) == "true" {
		query += " WHERE is_complete=TRUE"
	}

	if strings.ToLower(filter.Completed) == "false" {
		query += " WHERE is_complete=false"
	}

	if filter.Sort == "oldest" {
		query += " ORDER BY Created_at ASC"
	} else {
		query += " ORDER BY Created_at DESC"
	}

	rows, err := pgs.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []TodoModel

	for rows.Next() {
		var t TodoModel
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.IsComplete, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (pgs *PostgresRepository) ToggleTodo(ctx context.Context, id int) (TodoModel, error) {

	var updatedTodo TodoModel
	err := pgs.pool.QueryRow(ctx, "UPDATE todos SET is_complete= NOT is_complete, updated_at=$1 WHERE id=$2 RETURNING *", time.Now(), id).Scan(&updatedTodo.Id, &updatedTodo.Title, &updatedTodo.Description, &updatedTodo.IsComplete, &updatedTodo.CreatedAt, &updatedTodo.UpdatedAt)
	if err != nil {
		return TodoModel{}, err
	}
	return updatedTodo, err
}

func (pgs *PostgresRepository) UpdateTodo(ctx context.Context, id int, todo TodoModel) (TodoModel, error) {

	var updatedTodo TodoModel
	err := pgs.pool.QueryRow(ctx, "UPDATE todos SET title=$1, description=$2, is_complete=$3, updated_at=$4 WHERE id=$5 RETURNING *", todo.Title, todo.Description, todo.IsComplete, time.Now(), id).Scan(updatedTodo.Id, updatedTodo.Title, updatedTodo.Description, updatedTodo.IsComplete, updatedTodo.CreatedAt, updatedTodo.UpdatedAt)
	if err != nil {
		return TodoModel{}, err
	}
	return updatedTodo, nil

}

func (pgs *PostgresRepository) DeleteTodo(ctx context.Context, id int) error {

	_, err := pgs.pool.Exec(ctx, "DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
