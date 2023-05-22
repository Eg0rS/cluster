package errorhandler

import (
	"fmt"
	"log"
	"net/http"
)

func UnexpectedError(err error, res http.ResponseWriter) {
	log.Printf("Internal server error: %s", err)
	res.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(res, "%s", []byte("Internal server error"))
}

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "%s", []byte("Bad request"))
}

func BadParam(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "%s", []byte("Bad param"))
}

func BadBody(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "%s", []byte("Bad body"))
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%s", []byte("Not found"))
}

func Forbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, "%s", []byte("Forbidden"))
}

func AlreadyExist(w http.ResponseWriter) {
	w.WriteHeader(http.StatusConflict)
	fmt.Fprintf(w, "%s", []byte("Item already exist"))
}

func MethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "%s", []byte("Method Not Allowed"))
}

func Ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func CustomError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%s", []byte(err.Error()))
}
