package dbUtils

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

const pgxDriverName = "pgx"

func NewPostgres(connectionString string) (*sqlx.DB, error) {
	connConfig, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}
	if connConfig == nil {
		return nil, ConnConfigEmpty
	}

	nativeDB := stdlib.OpenDB(*connConfig)
	db := sqlx.NewDb(nativeDB, pgxDriverName)
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("can't ping to db: %s", err)
	}

	return db, nil
}
