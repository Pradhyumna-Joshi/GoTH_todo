package main

import (
	"log/slog"
	"os"

	"github.com/Pradhyumna-Joshi/go_todo_htmx/config"
)

func main() {

	connStr := "postgresql://test:db123@localhost:5432/todos"
	cfg := config.Config{
		Addr:    ":8080",
		ConnStr: connStr,
	}
	server := NewAPIServer(cfg)

	if err := server.run(server.mount()); err != nil {
		slog.Info("Failed to run server", "error", err)
		os.Exit(1)
	}
}
