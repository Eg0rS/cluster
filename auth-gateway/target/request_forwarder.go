package target

import "net/http"

type RequestForwarder interface {
	Forward(r *http.Request, payload string) (*http.Response, error)
}
