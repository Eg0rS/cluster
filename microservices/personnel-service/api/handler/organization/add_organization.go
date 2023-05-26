package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	organization "personnel_service/client/Organization"
	"personnel_service/lib"
	"personnel_service/service"
)

// AddOrganizationHandler
//
//		@Description	Add organization
//		@Tags			organization
//		@Accept			json
//		@Produce		json
//		@Param			request body organization.AddOrganizationRequest true "query params"
//		@Success		200	{object}	organization.AddOrganizationResponse
//	 	@Failure        400 {object}    organization.AddOrganizationBadResponse
//	 	@Failure        501 {object}    organization.AddOrganizationBadResponse
//		@Router			/new/organization [post]
func AddOrganizationHandler(logger *zap.SugaredLogger, service service.PersonnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request organization.AddOrganizationRequest
			res     []byte
		)

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			res, _ = json.Marshal(organization.AddOrganizationBadResponse{Id: 0, Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		id, err := service.AddOrganizations(r.Context(), request)
		if err != nil {
			logger.Debugf("Service error: %s", err)
			res, _ = json.Marshal(organization.AddOrganizationBadResponse{Id: 0, Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		res, _ = json.Marshal(organization.AddOrganizationResponse{Id: id, Error: ""})
		lib.SendResponse(w, http.StatusOK, res)
	}
}
