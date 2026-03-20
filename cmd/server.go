package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Pradhyumna-Joshi/go_todo_htmx/config"
	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/handler"
	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/repository"
	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type APIServer struct {
	config config.Config
}

func NewAPIServer(config config.Config) *APIServer {
	return &APIServer{config}
}

func (s *APIServer) mount() http.Handler {

	conn, err := pgxpool.New(context.Background(), s.config.ConnStr)
	if err != nil {
		slog.Info("Failed to Connect to Postgres", "error", err)
		return nil
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to TODOS"))
	})

	todoHandler := handler.NewTodoHandler(service.NewTodoService(repository.NewPostGresRepository(conn)))

	mux.HandleFunc("POST /todos", todoHandler.CreateTodo)
	mux.HandleFunc("GET /todos", todoHandler.GetTodos)
	mux.HandleFunc("PUT /todos/{id}", todoHandler.UpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", todoHandler.DeleteTodo)
	return mux
}

func (s *APIServer) run(h http.Handler) error {

	server := &http.Server{
		Addr:         s.config.Addr,
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}
	slog.Info("Server running!!", "port", s.config.Addr)
	return server.ListenAndServe()
}
