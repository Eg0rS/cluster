package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"personnel_service/api/handler"
	request "personnel_service/api/handler/request"
	"personnel_service/config"
	"personnel_service/lib/pctx"
	"personnel_service/service"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewServer(
	ctxProvider pctx.DefaultProvider,
	logger *zap.SugaredLogger,
	settings config.Settings,
	personnelService service.PersonnelService,
) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.PingHandle(logger)).Methods(http.MethodGet)

	router.HandleFunc("/personnel/new/radio_test", request.CreateRadioTestHandler(logger, personnelService)).Methods(http.MethodPost)
	router.HandleFunc("/personnel/new/text_test", request.CreateTextTest(logger, personnelService)).Methods(http.MethodPost)
	router.HandleFunc("/personnel/new/request", request.CreateRequestHandler(logger, personnelService)).Methods(http.MethodPost)
	router.HandleFunc("/personnel/get/requests/{user_id}", request.GetAllRequestsById(logger, personnelService)).Methods(http.MethodGet)
	router.HandleFunc("/personnel/get/test/{test_id}", request.GetTestsByTestIdHandler(logger, personnelService)).Methods(http.MethodGet)

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
