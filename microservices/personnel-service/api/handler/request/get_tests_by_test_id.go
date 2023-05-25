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

// GetTestsByTestIdHandler
//
//		@Description	Get test by him id
//		@Tags			personnel
//		@Accept			json
//		@Produce		json
//		@Param			test_id   path      int  true  "Test ID"
//		@Success		200	{object}	personnel.RadioTest
//	 	@Failure        400 {object}    personnel.CreateRequestResponse
//	 	@Failure        501 {object}    personnel.CreateRequestResponse
//		@Router			/personnel/get/test/{test_id} [get]
func GetTestsByTestIdHandler(logger *zap.SugaredLogger, personnelService service.PersonnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res []byte
		data, err := personnelService.GetTestByTestId(r.Context(), mux.Vars(r)["test_id"])
		if err != nil {
			logger.Debugf("Error on service: %s", err)
			res, _ = json.Marshal(personnel.CreateRequestResponse{Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		res, _ = json.Marshal(data)
		lib.SendResponse(w, http.StatusOK, res)
	}
}
