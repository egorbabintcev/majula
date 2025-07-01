package main

import (
	"log"
	"log/slog"
	"majula/internal/core"
	"majula/internal/infrastructure/storage/filesystem"
	"majula/internal/infrastructure/storage/inmem"
	web "majula/internal/interface/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	port := os.Getenv("MAJULA_PORT")
	fsStoragePath := os.Getenv("MAJULA_FS_STORAGE_PATH")

	if port == "" {
		port = "8000"
	}

	if fsStoragePath == "" {
		log.Fatal("fs storage path not specified")
	}

	ps := inmem.NewPackumentStorage()
	ts, err := filesystem.NewTarballStorage(l, fsStoragePath)

	if err != nil {
		log.Fatal(err)
	}

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
