package main

import (
	"log/slog"
	"majula/internal/core"
	"majula/internal/infrastructure"
	iface_http "majula/internal/interface/http"
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
	r := iface_http.NewRouter(s, l)

	http.ListenAndServe(addr, r)
}
