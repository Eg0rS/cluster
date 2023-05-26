package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"profile_service/api/handler"
	"profile_service/config"
	"profile_service/lib/pctx"
	"profile_service/service"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewServer(
	ctxProvider pctx.DefaultProvider,
	logger *zap.SugaredLogger,
	settings config.Settings,
	profileService service.ProfileService,
) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.PingHandle(logger)).Methods(http.MethodGet)

	router.HandleFunc("/update/{user_id}", handler.UpsertUserInfoHandler(logger, profileService)).Methods(http.MethodGet)
	router.HandleFunc("/get/info", handler.GetUserInfoHandler(logger, profileService)).Methods(http.MethodPost)

	// swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctxProvider()
		},
		Handler: router,
	}
}
