package authorization

import (
	"auth-gateway/config"
	"strings"
)

func NewAuthorizationHeader(value string, settings *config.Settings) *AuthorizationHeader {
	return &AuthorizationHeader{value, settings}
}

type AuthorizationHeader struct {
	Value    string
	settings *config.Settings
}

func (h *AuthorizationHeader) IsValid() bool {
	if h.Value == "" {
		return false
	}
	headerParts := strings.Split(h.Value, " ")
	if len(headerParts) != 2 {
		return false
	}
	if strings.ToLower(headerParts[0]) != "bearer" {
		return false
	}
	return true
}

func (h *AuthorizationHeader) GetToken() (*JWT, error) {
	headerParts := strings.Split(h.Value, " ")
	if len(headerParts) != 2 {
		return &JWT{}, nil
	}

	if strings.ToLower(headerParts[0]) != "bearer" {
		return &JWT{}, nil
	}

	token := strings.Split(h.Value, " ")[1]
	if token == "" {
		return &JWT{}, nil
	}
	return NewJWTToken(token, h.settings)
}
