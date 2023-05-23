package handlers

import (
	"auth/dal"
	"auth/utils/httpUtils"
	"net/http"

	"github.com/gorilla/mux"
)

type RefreshTokenHandler struct {
	RefreshToken dal.RefreshTokenRepository
}

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
