package handlers

import (
	"auth/dal"
	"log"
	"net/http"
)

// LogoutHandler
// @Description to_register user
// @Tags			auth
// @Accept			json
// @Produce		json
// @Router			/logout/ [post]
func LogoutHandler(refreshTokenRepository dal.RefreshTokenRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logout called")
	}
}
