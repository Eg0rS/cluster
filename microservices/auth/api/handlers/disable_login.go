package handlers

import (
	"auth/dal"
	"auth/utils/httpUtils"
	"encoding/json"
	"log"
	"net/http"
)

type DisableLoginRequest struct {
	UserId string
}

// DisableLoginHandler
// @Description to_register user
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			request body DisableLoginRequest true "query params"
// @Router			/token/disable-login/ [post]
func DisableLoginHandler(refreshTokenRepository dal.RefreshTokenRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := DisableLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err)
			httpUtils.BadRequest(w)
			return
		}
		err = refreshTokenRepository.DeleteByUserId(req.UserId)
		if err != nil {
			log.Println(err)
			httpUtils.BadRequest(w)
			return
		}

		return
	}
}
