package main

import (
	"log/slog"
	"net/http"

	"github.com/guilhermevialle/go-architecture/internal/user"
	"github.com/guilhermevialle/go-architecture/internal/user/infrastructure"
)

func main() {
	mux := http.NewServeMux()
	repo := infrastructure.NewInMemoryUserRepository()

	user.NewModule(mux, repo)

	addr := ":80"
	slog.Info("server listening", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("server error", "error", err)
	}
}
