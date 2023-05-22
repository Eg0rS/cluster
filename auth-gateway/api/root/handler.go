package root

import (
	"auth-gateway/api/root/authorization"
	"auth-gateway/config"
	"auth-gateway/microservices"
	"auth-gateway/target"
	"auth-gateway/utils/httpUtils"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func NewHandler(
	settings *config.Settings,
	requestForwarder target.RequestForwarder,
	authService *microservices.AuthService,
) *Handler {
	return &Handler{
		settings:         settings,
		requestForwarder: requestForwarder,
		authService:      authService,
	}
}

type Handler struct {
	settings         *config.Settings
	requestForwarder target.RequestForwarder
	authService      *microservices.AuthService
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var accessToken = &authorization.JWT{}
	var err error
	accessTokenString := authorization.GetToken(r)
	authHeader := authorization.NewAuthorizationHeader(
		accessTokenString,
		h.settings,
	)

	accessToken, err = authHeader.GetToken()
	if err != nil {
		// капец конечно костыли :-(
		accessToken = &authorization.JWT{}
	}

	serviceName := strings.TrimPrefix(r.URL.Path, "/")
	n := strings.Index(serviceName, "/")
	if n != -1 {
		serviceName = serviceName[:n]
	}

	payload := accessToken.GetPayload()

	resp, err := h.requestForwarder.Forward(r, payload)
	if err != nil {
		httpUtils.UnexpectedError(err, w)
		return
	}

	if resp == nil {
		httpUtils.UnexpectedError(err, w)
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	for headerName, headerValues := range resp.Header {
		for _, headerValue := range headerValues {
			w.Header().Add(headerName, headerValue)
		}
	}

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		httpUtils.UnexpectedError(err, w)
		return
	}
}

func isPrivateRoute(path, serviceName string) bool {
	return strings.HasPrefix(path, fmt.Sprintf("/%s/private", serviceName))
}
