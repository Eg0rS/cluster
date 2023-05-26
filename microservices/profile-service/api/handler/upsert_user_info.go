package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	profile "profile_service/client/Profile"
	"profile_service/lib"
	"profile_service/service"
	profileResponses "profile_service/swagger_responses"
)

// UpsertUserInfoHandler
//
//		@Description	Upsert user info
//		@Tags			profile
//		@Accept			json
//		@Produce		json
//		@Param			user_id   path      int  true  "User ID"
//		@Param			request body profile.UpsertUserInfoReq false "query params"
//		@Success		200	{object}	profile_responses.UpsertGoodResponse
//	 	@Failure        400 {object}    profile_responses.UpsertBadResponse
//	 	@Failure        501 {object}    profile_responses.UpsertBadResponse
//		@Router			/update/{user_id} [post]
func UpsertUserInfoHandler(logger *zap.SugaredLogger, service service.ProfileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request profile.UpsertUserInfoReq
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

		err = service.UpsertUserInfo(r.Context(), request, mux.Vars(r)["user_id"])
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			res, _ = json.Marshal(profileResponses.UpsertBadResponse{Error: err.Error()})
			lib.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		res, _ = json.Marshal(profileResponses.UpsertGoodResponse{})
		lib.SendResponse(w, http.StatusOK, res)
	}
}
