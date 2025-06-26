package main

import (
	"log/slog"
	"majula/internal/core"
	"majula/internal/infrastructure/storage/inmem"
	web "majula/internal/interface/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := os.Getenv("MAJULA_PORT")

	if port == "" {
		port = "8000"
	}

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	ps := inmem.NewPackumentStorage()
	ts := inmem.NewTarballStorage()
	s := core.NewService(ps, ts)
	srv := web.NewServer(s, l)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan

		srv.Stop()
	}()

	srv.Start(":" + port)
}
