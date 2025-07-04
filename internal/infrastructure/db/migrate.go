package db

import (
	"embed"
	"errors"

	"github.com/pressly/goose/v3"
)

//go:embed migrations
var migrations embed.FS

var (
	ErrApplyMigration = errors.New("failed to apply migrations")
)

func Migrate(c *Connection) error {
	goose.SetBaseFS(migrations)

	err := goose.SetDialect("sqlite")

	if err != nil {
		return errors.Join(ErrApplyMigration, err)
	}

	err = goose.Up(c.DB, "migrations")

	if err != nil {
		return errors.Join(ErrApplyMigration, err)
	}

	return nil
}
