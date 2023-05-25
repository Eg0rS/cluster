package handle

import (
	"encoding/json"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	personnel "personnel_service/client/Personnel"
	"personnel_service/lib"
	"personnel_service/service"
)

//	    CreateTextTest
//
//		@Description	Create text test and return test id
//		@Tags			personnel
//		@Accept multipart/form-data
//		@Produce		json
//		@Param file formData file true "file"
//		@Param title formData string true "title"
//		@Param description formData string true "description"
//		@Success		200	{object}	swagger_responses.CreateTestOkRes
//		@Failure        400 {object}    swagger_responses.HTTPErrorCreateTest
//		@Failure        501 {object}    swagger_responses.HTTPErrorCreateTest
//		@Router	/personnel/new/text_test [post]
func CreateTextTest(logger *zap.SugaredLogger, personnelService service.PersonnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res []byte

		testTitle := r.FormValue("title")
		testDescription := r.FormValue("description")
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			logger.Debugf("File open error: %s", err)
			res, _ = json.Marshal(personnel.CreateTestResponse{Id: 0, Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				logger.Debugf("File close error: %s", err)
				res, _ = json.Marshal(personnel.CreateTestResponse{Id: 0, Error: err.Error()})
				lib.SendResponse(w, http.StatusBadRequest, res)
				return
			}
		}(file)

		model := personnel.TextTest{
			Title:       testTitle,
			Description: testDescription,
			File:        file,
			Header:      fileHeader,
		}

		id, err := personnelService.CreateTestText(r.Context(), model)
		if err != nil {
			logger.Debugf("Service error: %s", err)
			res, _ = json.Marshal(personnel.CreateTestResponse{Id: 0, Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		res, _ = json.Marshal(personnel.CreateTestResponse{Id: id, Error: ""})
		lib.SendResponse(w, http.StatusOK, res)
	}
}
