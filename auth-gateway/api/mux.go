package api

import (
	"auth-gateway/api/root"
	"net/http"
)

func NewMux(rootHandler *root.Handler) *http.ServeMux {

	mux := http.NewServeMux()
	mux.Handle("/", rootHandler)
	return mux
}
