package handle

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	personnel "personnel_service/client/Personnel"
	"personnel_service/lib"
	"personnel_service/service"
)

// GetAllRequestsById
//
//		@Description	Get all requests by user id
//		@Tags			personnel
//		@Accept			json
//		@Produce		json
//		@Param			user_id   path      int  true  "User ID"
//		@Success		200	{object}	swagger_responses.GetRequestsOkRes
//	 	@Failure        400 {object}    swagger_responses.HTTPErrorGetRequests
//	 	@Failure        501 {object}    swagger_responses.HTTPErrorGetRequests
//		@Router			/personnel/get/requests/{user_id} [get]
func GetAllRequestsById(logger *zap.SugaredLogger, personnelService service.PersonnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res []byte
		)

		requests, err := personnelService.GetAllRequestsByUserId(r.Context(), mux.Vars(r)["user_id"])
		if err != nil {
			logger.Debugf("Error on create request: %s", err)
			res, _ = json.Marshal(personnel.GetRequestsResponse{Error: err.Error(), Requests: requests})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		res, _ = json.Marshal(personnel.GetRequestsResponse{Requests: requests, Error: ""})
		lib.SendResponse(w, http.StatusOK, res)
	}
}
