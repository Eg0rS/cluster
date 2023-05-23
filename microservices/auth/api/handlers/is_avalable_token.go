package handlers

import (
	"auth/dal"
	"auth/utils/httpUtils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type CheckTokenRequest struct {
	Token string `json:"token"`
}

type CheckTokenResponse struct {
	Exists bool `json:"exists"`
}

func IsAvailableToken(refreshTokenRepository dal.RefreshTokenRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqModel := CheckTokenRequest{}
		err := json.NewDecoder(r.Body).Decode(&reqModel)
		if err != nil {
			log.Println(err)
			httpUtils.BadRequest(w)
			return
		}

		exist := refreshTokenRepository.AccessTokenExists(strings.TrimPrefix(reqModel.Token, "Bearer "))

		respModel := CheckTokenResponse{}
		respModel.Exists = exist
		err = json.NewEncoder(w).Encode(respModel)
		if err != nil {
			log.Println(err)
			httpUtils.BadRequest(w)
			return
		}
	}
}
