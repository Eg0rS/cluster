package root

import (
	"io"
	"net/http"
)

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if _, err := io.WriteString(w, "pong"); err != nil {

		}
	}
}
