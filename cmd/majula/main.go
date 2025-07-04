package main

import (
	"log"
	"log/slog"
	"majula/internal/core"
	"majula/internal/infrastructure/db"
	"majula/internal/infrastructure/storage/inmem"
	"majula/internal/infrastructure/tarball"
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
	dbConnectionString := os.Getenv("MAJULA_DB_CONNECTION_STRING")

	if port == "" {
		port = "8000"
	}

	if fsStoragePath == "" {
		fsStoragePath = "/var/majula"
	}

	if dbConnectionString == "" {
		dbConnectionString = "file:/var/majula/db.sqlite?cache=shared&mode=rwc"
	}

	conn, err := db.NewConnection(l, dbConnectionString)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	ps := inmem.NewPackageStorage()
	ts, err := tarball.NewStorage(l, fsStoragePath)

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
		conn.Close()
	}()

	srv.Start(":" + port)
}
