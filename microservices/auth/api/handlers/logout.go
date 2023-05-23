package handlers

import (
	"auth/dal"
	"log"
	"net/http"
)

func LogoutHandler(refreshTokenRepository dal.RefreshTokenRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logout called")
	}
}
