package authorize

import (
	"auth-gateway/lib/error-handler"
	"context"
	"net/http"
)

// NewAuthHandler Позволяет обернуть обработчик запроса для проверки авторизации. В случае успешной проверки задает в Context
// переменную "auth" типа authorize.Authorize
func NewAuthHandler(origin http.Handler) http.Handler {
	return handler{
		origin: origin,
	}
}

type handler struct {
	origin http.Handler
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth, err := GetAuth(r)
	if err != nil {
		errorhandler.Forbidden(w)
		return
	}
	h.origin.ServeHTTP(w, r.WithContext(newAuthContext(r.Context(), auth)))
}

func newAuthContext(ctx context.Context, auth Authorize) context.Context {
	return context.WithValue(ctx, "auth", auth)
}
