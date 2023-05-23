package main

import (
	"auth/api"
	"auth/api/authentication"
	"auth/api/authentication/generation"
	"auth/api/handlers"
	"auth/config"
	"auth/dal"
	"auth/microservices"
	"auth/microservices/accountservice"
	"auth/tokencache"
	"context"
)

type serviceProvider struct {
	settings               *config.Settings
	userRepository         dal.UserRepository
	userRoleRepository     dal.UserRoleRepository
	refreshTokenRepository dal.RefreshTokenRepository
	passwordHasher         microservices.PasswordHasherService
	done                   chan struct{}
	mainCtx                context.Context
}

func (p serviceProvider) provide() *api.Server {
	accessTokenGenerator := generation.NewAccessTokenGenerator(
		p.settings,
		generation.NewRoleChecker(p.userRoleRepository),
		accountservice.NewService(p.settings.MicroserviceAccountService),
	)
	refreshTokenGenerator := generation.NewRefreshTokenGenerator(p.settings)
	authenticator := authentication.NewAuthenticator(
		p.userRepository,
		accessTokenGenerator,
		refreshTokenGenerator,
		p.refreshTokenRepository,
		p.passwordHasher,
		p.settings,
		p.clickRepository,
		p.authLog,
	)
	tokenRequestHandler := handlers.NewTokenRequestHandler(
		authenticator,
		authentication.NewAuthenticatorByPasswordHash(
			p.userRepository,
			accessTokenGenerator,
			refreshTokenGenerator,
			p.refreshTokenRepository,
			p.passwordHasher,
			p.settings,
		),
		authentication.NewRefresher(
			p.userRepository,
			accessTokenGenerator,
			refreshTokenGenerator,
			generation.NewRefreshTokenParser(p.settings),
			p.refreshTokenRepository,
			p.settings),
		authentication.NewAuthenticatorByUUID(
			p.settings,
			authenticator,
			p.uuidRepository,
			p.userRepository),
		p.loggerSlack,
		p.availableUsersRepo,
	)

	tokenCache := tokencache.New("", "")

	return api.NewAuthServer(
		p.settings,
		p.done,
		api.NewMux(
			tokenRequestHandler,
			p.refreshTokenRepository,
			p.loggerService,
			p.mainCtx,
			tokenCache),
	)
}
