package main

import (
	"context"
	"net/http"
	"personnel_service/api"
	"personnel_service/config"
	"personnel_service/database"
	"personnel_service/database/pesrsonnel_repo"
	"personnel_service/lib/pctx"
	"personnel_service/service"

	"go.uber.org/zap"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *http.Server
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {
	pgDb, err := database.NewPgx(settings.Postgres)
	if err != nil {
		panic(err)
	}

	var (
		personnelRepo = pesrsonnel_repo.NewPersonnelRepository(logger, pgDb)

		personnelService = service.NewPersonnelService(logger, personnelRepo)

		server = api.NewServer(ctxProvider, logger, settings, personnelService)
	)

	return App{
		logger:   logger,
		settings: settings,
		server:   server,
	}
}

func (a App) Run() {
	go func() {
		_ = a.server.ListenAndServe()
	}()
	a.logger.Debugf("HTTP server started on %d", a.settings.Port)
}

func (a App) Stop(ctx context.Context) {
	_ = a.server.Shutdown(ctx)
	a.logger.Debugf("HTTP server stopped")
}
