package main

import (
	"log/slog"
	"net/http"

	"github.com/guilhermevialle/go-architecture/internal/order"
	"github.com/guilhermevialle/go-architecture/internal/user"
)

func main() {
	mux := http.NewServeMux()

	user.NewModule(mux)
	order.NewModule(mux)

	addr := ":80"
	slog.Info("server listening", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("server error", "error", err)
	}
}
