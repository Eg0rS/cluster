package handle

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	personnel "personnel_service/client/Personnel"
	"personnel_service/lib"
	"personnel_service/service"
)

// CreateRadioTestHandler
//
//		@Description	Create radio test and return test id
//		@Tags			personnel
//		@Accept			json
//		@Produce		json
//		@Param			request body personnel.RadioTest false "query params"
//		@Success		200	{object}	swagger_responses.CreateTestOkRes
//	 	@Failure        400 {object}    swagger_responses.HTTPErrorCreateTest
//	 	@Failure        501 {object}    swagger_responses.HTTPErrorCreateTest
//		@Router			/personnel/new/radio_test [post]
func CreateRadioTestHandler(logger *zap.SugaredLogger, personnelService service.PersonnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request personnel.RadioTest
			res     []byte
		)

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			res, _ = json.Marshal(personnel.CreateTestResponse{Id: 0, Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		id, err := personnelService.CreateRadioTest(r.Context(), request)
		if err != nil {
			logger.Debugf("Error on create request: %s", err)
			res, _ = json.Marshal(personnel.CreateTestResponse{Id: 0, Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		res, _ = json.Marshal(personnel.CreateTestResponse{Id: id, Error: ""})
		lib.SendResponse(w, http.StatusOK, res)
	}
}
