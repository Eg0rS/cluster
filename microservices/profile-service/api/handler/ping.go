package handler

import (
	"go.uber.org/zap"
	"io"
	"net/http"
)

// PingHandle
// @Description ping-pong ops...
// @Router /ping [get]
func PingHandle(logger *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "pong")
		if err != nil {
			logger.Debugf("Error: %s", err)
		}
	}
}
