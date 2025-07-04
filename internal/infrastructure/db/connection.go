package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	_ "modernc.org/sqlite"
)

var (
	ErrOpenConnection = errors.New("failed to open database connection")
)

type Connection struct {
	logger *slog.Logger
	DB     *sql.DB
}

func NewConnection(l *slog.Logger, dsn string) (*Connection, error) {
	l = l.With(slog.String("component", "database"))

	l.Info("Open database connection")
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		return nil, errors.Join(ErrOpenConnection, err)
	}

	return &Connection{
		logger: l,
		DB:     db,
	}, nil
}

func (c *Connection) Close() {
	err := c.DB.Close()

	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close database connection %s", err))
	}
}
