package api

import (
	"auth/api/handlers"
	"auth/dal"
	"github.com/gorilla/mux"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewMux(
	tokenRequestHandler *handlers.TokenRequestHandler,
	refreshTokenRepository dal.RefreshTokenRepository,
	userRepository dal.UserRepository,
) http.Handler {

	router := mux.NewRouter()

	router.Handle("/token", jsonHttpHandler{method: tokenRequestHandler.HandleTokenRequest})
	router.Handle("/register", handlers.Register(userRepository)).Methods(http.MethodPost)
	router.Handle("/token/{userId}", &handlers.RefreshTokenHandler{RefreshToken: refreshTokenRepository}).Methods(http.MethodDelete)
	router.Handle("/token-exist/", handlers.IsAvailableToken(refreshTokenRepository)).Methods(http.MethodPost)
	router.Handle("/token/disable-login/", handlers.DisableLoginHandler(refreshTokenRepository)).Methods(http.MethodPost)
	router.Handle("/logout/", handlers.LogoutHandler(refreshTokenRepository)).Methods(http.MethodPost)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
