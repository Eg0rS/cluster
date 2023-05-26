package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"personnel_service/api/handler"
	organization "personnel_service/api/handler/organization"
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

	router.HandleFunc("/new/radio_test", request.CreateRadioTestHandler(logger, personnelService)).Methods(http.MethodPost)
	router.HandleFunc("/new/text_test", request.CreateTextTest(logger, personnelService)).Methods(http.MethodPost)
	router.HandleFunc("/new/request", request.CreateRequestHandler(logger, personnelService)).Methods(http.MethodPost)
	router.HandleFunc("/get/requests/{user_id}", request.GetAllRequestsById(logger, personnelService)).Methods(http.MethodGet)
	router.HandleFunc("/get/test/{test_id}", request.GetTestsByTestIdHandler(logger, personnelService)).Methods(http.MethodGet)

	router.HandleFunc("/new/organization", organization.AddOrganizationHandler(logger, personnelService)).Methods(http.MethodPost)
	router.HandleFunc("/get/organizations", organization.GetAllOrganizationsHandler(logger, personnelService)).Methods(http.MethodGet)

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
