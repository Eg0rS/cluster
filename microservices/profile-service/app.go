package main

import (
	"context"
	"net/http"
	"profile_service/api"
	"profile_service/config"
	"profile_service/database"
	"profile_service/database/profile_repo"
	"profile_service/lib/pctx"
	"profile_service/service"

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
		profileRepo = profile_repo.NewProfileRepository(logger, pgDb)

		profileService = service.NewProfileService(logger, profileRepo)

		server = api.NewServer(ctxProvider, logger, settings, profileService)
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
