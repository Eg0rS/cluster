package handlers

import (
	"auth/dal"
	"auth/utils/httpUtils"
	"github.com/gorilla/mux"
	"net/http"
)

type RefreshTokenHandler struct {
	RefreshToken dal.RefreshTokenRepository
}

// DeleteByUserId
// @Description delete refresh token by user id user
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			user_id   path      int  true  "User ID"
// @Router			/token/{userId} [delete]
func (h *RefreshTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId, ok := mux.Vars(r)["userId"]

	if !ok || len(userId) == 0 {
		httpUtils.BadRequest(w)
		return
	}

	err := h.RefreshToken.DeleteByUserId(userId)

	if err != nil {
		httpUtils.BadRequest(w)
		return
	}
}
