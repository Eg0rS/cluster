package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	profile "profile_service/client/Profile"
	"profile_service/lib"
	"profile_service/service"
	profileResponses "profile_service/swagger_responses"
)

// GetUserInfoHandler
//
//		@Description	Get user info
//		@Tags			profile
//		@Accept			json
//		@Produce		json
//		@Param			request body profile.GetUserInfoReq true "query params"
//		@Success		200	{object}	profile.UpsertUserInfoReq
//	 	@Failure        400 {object}    profile_responses.UpsertBadResponse
//	 	@Failure        501 {object}    profile_responses.UpsertBadResponse
//		@Router			/get/info [post]
func GetUserInfoHandler(logger *zap.SugaredLogger, service service.ProfileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request profile.GetUserInfoReq
			res     []byte
		)

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			res, _ = json.Marshal(profileResponses.UpsertBadResponse{Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		data, err := service.GetUserInfo(r.Context(), request.RefreshToken)
		if err != nil {
			logger.Debugf("Service error: %s", err)
			res, _ = json.Marshal(profileResponses.UpsertBadResponse{Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}
		response := profile.MapServiceUpsertInfoToClient(data)
		res, _ = json.Marshal(response)
		lib.SendResponse(w, http.StatusOK, res)
	}
}
