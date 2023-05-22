package microservices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AuthService struct {
	url string
}

type CheckTokenRequest struct {
	Token string `json:"token"`
}

type CheckTokenResponse struct {
	Exists bool `json:"exists"`
}

func NewAuthService(url string) *AuthService {
	return &AuthService{
		url: url,
	}
}

func (s *AuthService) IsAccessTokenExists(accessToken string) bool {
	request := CheckTokenRequest{
		Token: accessToken,
	}
	b := &bytes.Buffer{}
	err := json.NewEncoder(b).Encode(request)
	if err != nil {
		log.Println(err)
		return false
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/token-exist/", s.url), b)
	if err != nil {
		log.Println(err)
		return false
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	if resp.Body == nil || resp.StatusCode != 200 {
		log.Printf("body is nil or response not 200 but %d\n", resp.StatusCode)
		return false
	}
	response := CheckTokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return false
	}

	return response.Exists
}
