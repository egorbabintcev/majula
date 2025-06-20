package main

import (
	"log/slog"
	"majula/internal/core"
	"majula/internal/infrastructure"
	web "majula/internal/interface/http"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("MAJULA_REGISTRY_PORT")

	if port == "" {
		port = "8000"
	}

	addr := ":" + port

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	st := infrastructure.NewMemoryStorage()
	s := core.NewService(st)
	r := web.NewRouter(s, l)

	http.ListenAndServe(addr, r)
}
