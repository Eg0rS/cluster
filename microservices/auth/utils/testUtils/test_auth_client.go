package testUtils

import (
	"auth/config"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func NewAuthClient(settings *config.Settings, userAgent string, body string) *TestAuthClient {
	return &TestAuthClient{
		host:       "http://localhost:" + strconv.Itoa(settings.Port),
		httpClient: &http.Client{},
		userAgent:  userAgent,
		body:       body,
	}
}

type TestAuthClient struct {
	host         string
	clientSecret string
	httpClient   *http.Client
	userAgent    string
	body         string
}

func (c *TestAuthClient) PerformRequest() (string, error) {
	req, _ := http.NewRequest("POST", c.host+"/token", bytes.NewReader([]byte(c.body)))
	req.Header.Add("User-Agent", c.userAgent)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == 400 {
		return "", BadRequest
	}
	if res.StatusCode == 401 {
		return "", Unauthorized
	}
	if res.StatusCode == 403 {
		return "", Forbidden
	}
	if res.StatusCode >= 500 {
		return "", InternalServer
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("unrecognize error")
	}
	body, err := ioutil.ReadAll(res.Body)
	return string(body), err
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	UserAgent string `json:"user_agent"`
}

const (
	BadRequest     = Error("400")
	Unauthorized   = Error("401")
	Forbidden      = Error("403")
	InternalServer = Error("500")
)

type Error string

func (e Error) Error() string { return string(e) }
