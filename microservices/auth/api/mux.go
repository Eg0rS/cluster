package api

import (
	"auth/api/handlers"
	"auth/dal"
	"auth/microservices"
	"auth/tokencache"
	"context"
	"log"
	"net/http"

	slHttp "gitea.gospodaprogrammisty.ru/Go/servicelib/logging/http"
	"github.com/gorilla/mux"
)

func NewMux(
	tokenRequestHandler *handlers.TokenRequestHandler,
	refreshTokenRepository dal.RefreshTokenRepository,
	loggerService microservices.LoggerService,
	mainCtx context.Context,
	tokenCache *tokencache.TokenCache,
) http.Handler {
	httpLogger := getHttpLogger(mainCtx, loggerService)

	router := mux.NewRouter()

	router.Handle("/token", httpLogger.LogHTTP(jsonHttpHandler{method: tokenRequestHandler.HandleTokenRequest}))
	router.Handle("/token/{userId}", &handlers.RefreshTokenHandler{RefreshToken: refreshTokenRepository}).Methods(http.MethodDelete)
	router.Handle("/token-exist/", handlers.IsAvailableToken(refreshTokenRepository)).Methods(http.MethodPost)
	router.Handle("/token/disable-login/", handlers.DisableLoginHandler(refreshTokenRepository, tokenCache)).Methods(http.MethodPost)
	router.Handle("/logout/", handlers.LogoutHandler(refreshTokenRepository)).Methods(http.MethodPost)

	return router
}

func getHttpLogger(ctx context.Context, logger microservices.LoggerService) slHttp.Logger {
	return slHttp.Logger{LogFunc: func(entry slHttp.Entry) {
		err := logger.LogAccess(ctx, entry)
		if err != nil {
			log.Println(err.Error())
		}
	}}
}
