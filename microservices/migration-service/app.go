package main

import (
	"context"
	"migration-service/config"
	"migration-service/database"
)

type App struct {
}

func NewApp(settings config.Settings) App {
	pgDb, err := database.NewPgx(settings.Postgres)
	if err != nil {
		panic(err)
	}

	err = database.ResetMigrations(pgDb)
	if err != nil {
		panic(err)
	}

	err = database.UpMigrations(pgDb)
	if err != nil {
		panic(err)
	}

	return App{}
}

func (a App) Run() {

}

func (a App) Stop(ctx context.Context) {

}
