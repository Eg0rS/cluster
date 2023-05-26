package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

func UpMigrations(db *sqlx.DB) error {
	return goose.Up(db.DB, "./database/migrations/")
}

func ResetMigrations(db *sqlx.DB) error {
	return goose.Reset(db.DB, "./database/migrations/")
}
