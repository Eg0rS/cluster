package handlers

import (
	"auth/api/authentication"
	"encoding/json"
	"log"
	"net/http"
)

func NewTokenRequestHandler(
	authenticator *authentication.Authenticator,
	authenticatorByPasswordHash *authentication.AuthenticatorByPasswordHash,
	refresher *authentication.Refresher,
) *TokenRequestHandler {
	return &TokenRequestHandler{
		authenticator:               authenticator,
		authenticatorByPasswordHash: authenticatorByPasswordHash,
		refresher:                   refresher,
	}
}

type TokenRequestHandler struct {
	authenticator               *authentication.Authenticator
	authenticatorByPasswordHash *authentication.AuthenticatorByPasswordHash
	refresher                   *authentication.Refresher
}

// HandleTokenRequest
// @Description to_auth user
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			request body TokenRequest true "query params"
// @Success		200	{object}	TokenResponse
// @Router			/token [post]
func (h *TokenRequestHandler) HandleTokenRequest(request *http.Request) (interface{}, error) {
	requestData, err := h.parseRequestData(request)
	if err != nil {
		return nil, err
	}
	switch requestData.GrantType {
	case PasswordGrantType:
		log.Println("Authorize by " + PasswordGrantType)
		tokenResponse, err := h.authByCredentials(requestData)
		if err != nil {
			return nil, err
		}
		return tokenResponse, nil
	case RefreshTokenGrantType:
		log.Println("Authorize by " + RefreshTokenGrantType)
		tokenResponse, err := h.authByRefreshToken(requestData)
		if err != nil {
			return nil, err
		}
		return tokenResponse, nil
	default:
		return nil, BadRequestError
	}
}

func (h *TokenRequestHandler) parseRequestData(request *http.Request) (*TokenRequest, error) {
	if request.Method != http.MethodPost {
		return nil, BadRequestError
	}
	decoder := json.NewDecoder(request.Body)
	var requestData TokenRequest
	err := decoder.Decode(&requestData)
	if err != nil {
		return nil, BadRequestError
	}
	defer request.Body.Close()
	return &requestData, nil
}

func (h *TokenRequestHandler) authByCredentials(request *TokenRequest) (*TokenResponse, error) {
	authorizeRequest := authentication.AuthorizeRequest{
		Credentials: authentication.Credentials{
			Email:    request.UserName,
			Password: request.Password,
		},
	}

	access, _, err := h.authenticator.Authenticate(&authorizeRequest)
	if err != nil {
		log.Println(err)
	}
	log.Println("Authenticate proshla uspeshno")
	if err != nil {
		switch err {

		case authentication.InvalidClientSecretError:
			return nil, ForbiddenError
		case authentication.InvalidCredentialsError:

			return nil, ForbiddenError
		default:
			return nil, err
		}
	}

	return &TokenResponse{
		AccessToken:  access.AccessToken.Token,
		ExpiresIn:    access.AccessToken.TTL,
		RefreshToken: access.RefreshToken.Token,
		TokenType:    "bearer",
	}, nil
}

func (h *TokenRequestHandler) authByRefreshToken(request *TokenRequest) (*TokenResponse, error) {
	refreshRequest := authentication.RefreshRequest{
		RefreshToken: request.RefreshToken,
	}
	access, err := h.refresher.RefreshAccess(&refreshRequest)
	if err != nil {
		log.Println(err)
		switch err {
		case authentication.InvalidClientSecretError:
			return nil, ForbiddenError
		case authentication.InvalidTokenError:
			return nil, ForbiddenError
		default:
			return nil, err
		}
	}

	return &TokenResponse{
		AccessToken:  access.AccessToken.Token,
		ExpiresIn:    access.AccessToken.TTL,
		RefreshToken: access.RefreshToken.Token,
		TokenType:    "bearer",
	}, nil
}

func (h *TokenRequestHandler) authByUserNameAndPasswordHash(request *TokenRequest) (*TokenResponse, error) {
	authorizeRequest := authentication.AuthorizeRequest{
		Credentials: authentication.Credentials{
			Email:    request.UserName,
			Password: request.Password,
		},
	}

	access, err := h.authenticatorByPasswordHash.Authenticate(&authorizeRequest)
	if err != nil {
		switch err {

		case authentication.InvalidClientSecretError:
			return nil, ForbiddenError
		case authentication.InvalidCredentialsError:
			return nil, ForbiddenError
		default:
			return nil, err
		}
	}

	return &TokenResponse{
		AccessToken:  access.AccessToken.Token,
		ExpiresIn:    access.AccessToken.TTL,
		RefreshToken: access.RefreshToken.Token,
		TokenType:    "bearer",
	}, nil
}
