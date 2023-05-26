package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"personnel_service/lib"
	"personnel_service/model"
	"personnel_service/service"
)

// GetAllOrganizationsHandler
//
//		@Description	Get all organization
//		@Tags			organization
//		@Produce		json
//		@Success		200	{object}	organization.OrganizationsInfoResponse
//	 	@Failure        400 {object}    model.GetOrganizationsModel
//	 	@Failure        501 {object}    model.GetOrganizationsModel
//		@Router			/get/organizations [get]
func GetAllOrganizationsHandler(logger *zap.SugaredLogger, service service.PersonnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res []byte

		org, err := service.GetOrganizations(r.Context())
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			res, _ = json.Marshal(model.GetOrganizationsModel{})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		res, _ = json.Marshal(org)
		lib.SendResponse(w, http.StatusOK, res)
	}
}
