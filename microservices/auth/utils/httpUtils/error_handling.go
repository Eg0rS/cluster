package httpUtils

import (
	"fmt"
	"log"
	"net/http"
)

func UnexpectedError(err error, res http.ResponseWriter) {
	log.Printf("Internal server error: %s\n", err)
	res.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(res, "%s", []byte("Internal server error"))
}

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "%s", []byte("Bad request"))
}

func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "%s", []byte("Unauthorized"))
}

func Forbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, "%s", []byte("Forbidden"))
}
