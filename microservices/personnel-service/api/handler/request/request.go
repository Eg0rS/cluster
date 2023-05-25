package handle

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	personnel "personnel_service/client/Personnel"
	"personnel_service/lib"
	"personnel_service/service"
)

// CreateRequestHandler
//
//		@Description	Create request
//		@Tags			personnel
//		@Accept			json
//		@Produce		json
//		@Param			request body personnel.Request false "query params"
//		@Success		200	{object}	swagger_responses.CreateRequestOkRes
//	 	@Failure        400 {object}    swagger_responses.HTTPErrorCreateRequest
//	 	@Failure        501 {object}    swagger_responses.HTTPErrorCreateRequest
//		@Router			/personnel/new/request [post]
func CreateRequestHandler(logger *zap.SugaredLogger, personnelService service.PersonnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request personnel.Request
			res     []byte
		)

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			res, _ = json.Marshal(personnel.CreateRequestResponse{Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		err = personnelService.CreateRequest(r.Context(), request)
		if err != nil {
			logger.Debugf("Error on create request: %s", err)
			res, _ = json.Marshal(personnel.CreateRequestResponse{Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		res, _ = json.Marshal(personnel.CreateRequestResponse{Error: ""})
		lib.SendResponse(w, http.StatusOK, res)
	}
}
