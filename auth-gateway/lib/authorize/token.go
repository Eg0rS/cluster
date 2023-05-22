package authorize

import (
	"auth-gateway/errors"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

type Authorize struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Iat      int    `json:"iat"`
	Exp      int    `json:"exp"`
}

func GetToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func GetAuth(r *http.Request) (auth Authorize, err error) {
	return Parse(GetToken(r))
}

func Parse(token string) (auth Authorize, err error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return Authorize{}, errors.New("wrong authorization parts")
	}

	return decode(parts[1])
}

func decode(token string) (Authorize, error) {
	buf, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return Authorize{}, err
	}

	auth := Authorize{}
	err = json.Unmarshal(buf, &auth)
	if err != nil {
		return Authorize{}, err
	}

	return auth, nil
}

func GetAuthByContext(ctx context.Context) (auth *Authorize, err error) {
	if authContext, ok := ctx.Value("auth").(Authorize); !ok {
		return nil, errors.New("there is no authorization")
	} else {
		return &authContext, nil
	}
}
