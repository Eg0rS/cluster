package main

import (
	"auth-gateway/api"
	"auth-gateway/api/root"
	"auth-gateway/config"
	"auth-gateway/microservices"
	"auth-gateway/target"
)

type serviceProvider struct {
	settings      *config.Settings
	serverStopped chan struct{}
}

func (sl serviceProvider) provideServer() *api.Server {
	authService := microservices.NewAuthService(sl.settings.AuthService)

	return api.NewServer(
		sl.settings,
		api.NewMux(
			root.NewHandler(
				sl.settings,
				target.NewHTTPRequestForwarder(sl.settings),
				authService,
			),
		),
		sl.serverStopped,
	)
}
