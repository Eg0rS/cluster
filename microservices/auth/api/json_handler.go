package api

import (
	"auth/api/authentication"
	"auth/api/handlers"
	"auth/utils/httpUtils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type jsonHttpHandler struct {
	method func(*http.Request) (interface{}, error)
}

func (h jsonHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resultObject, err := h.method(r)
	if err != nil {
		handleError(err, w)
		return
	}

	resultJson, err := json.Marshal(resultObject)
	if err != nil {
		httpUtils.UnexpectedError(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s", resultJson)
}

func handleError(err error, w http.ResponseWriter) {
	if userIsNotApprovedErr, ok := err.(authentication.UserIsNotApprovedError); ok {
		errJson, jsonErr := json.Marshal(userIsNotApprovedErr)
		if jsonErr != nil {
			httpUtils.UnexpectedError(jsonErr, w)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		_, errWrite := w.Write(errJson)
		if errWrite != nil {
			log.Printf("не удалось записать ответ в responseWriter: %s", errWrite)
			httpUtils.UnexpectedError(errWrite, w)
			return
		}
		return
	}
	zn, ok := err.(authentication.TimeOutAttemptErr)
	if ok {
		_, err := fmt.Fprintf(w, "%s", []byte(`{"TimeOutAttempt":"`+zn.TimeOut+`"}`))
		if err != nil {
			log.Println(err)
		}
		return
	}
	switch err {
	case handlers.BadRequestError:
		httpUtils.BadRequest(w)
		return
	case handlers.UnauthorizedError:

		httpUtils.Unauthorized(w)
		return
	case handlers.ForbiddenError:
		httpUtils.Forbidden(w)
		return
	default:
		httpUtils.UnexpectedError(err, w)
		return
	}
}
