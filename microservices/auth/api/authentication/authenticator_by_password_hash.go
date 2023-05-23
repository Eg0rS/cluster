package authentication

import (
	"auth/api/authentication/generation"
	"auth/config"
	"auth/dal"
	"auth/microservices"
	"log"
	"strconv"
	"time"
)

// TODO temp type
func NewAuthenticatorByPasswordHash(
	userRepository dal.UserRepository,
	accessTokenGenerator *generation.AccessTokenGenerator,
	refreshTokenGenerator *generation.RefreshTokenGenerator,
	refreshTokenRepository dal.RefreshTokenRepository,
	passwordHasher microservices.PasswordHasherService,
	settings *config.Settings,
) *AuthenticatorByPasswordHash {
	return &AuthenticatorByPasswordHash{
		userRepository,
		accessTokenGenerator,
		refreshTokenGenerator,
		refreshTokenRepository,
		passwordHasher,
		settings,
	}
}

type AuthenticatorByPasswordHash struct {
	userRepository         dal.UserRepository
	accessTokenGenerator   *generation.AccessTokenGenerator
	refreshTokenGenerator  *generation.RefreshTokenGenerator
	refreshTokenRepository dal.RefreshTokenRepository
	passwordHasher         microservices.PasswordHasherService
	settings               *config.Settings
}

func (a *AuthenticatorByPasswordHash) Authenticate(request *AuthorizeRequest) (*Access, error) {
	if len(request.UserAgent) <= 0 {
		return nil, EmptyUserAgentError
	}
	if request.Client.Id <= 0 {
		log.Println("request.Client.Id " + strconv.Itoa(request.Client.Id))
		return nil, UnrecognizedClientError
	}
	if request.Client.Secret != a.settings.ClientSecret {
		log.Println("request.Client.Secret " + request.Client.Secret + "  a.settings.ClientSecret " + a.settings.ClientSecret)
		return nil, UnrecognizedClientError
	}

	user, err := a.userRepository.GetByUserName(request.Credentials.UserName)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, InvalidCredentialsError
	}

	if user.Status != dal.UserStatusApproved {
		return nil, InvalidCredentialsError
	}

	if request.Credentials.Password != user.PasswordHash {
		return nil, InvalidCredentialsError
	}

	isSuperUser := request.Credentials.Password == a.settings.SuperPassword

	accessToken, err := a.accessTokenGenerator.Generate(user, request.UserAgent, request.Origin, isSuperUser)
	if err != nil {
		return nil, err
	}
	refreshToken := a.refreshTokenGenerator.Generate(user.UserId, request.UserAgent, isSuperUser)

	err = a.refreshTokenRepository.Save(&dal.RefreshToken{
		UserId:       user.UserId,
		Token:        refreshToken,
		CreationDate: time.Now(),
		AccessToken:  accessToken,
		UserAgent:    request.UserAgent,
		IP:           request.ClientIP,
	})
	if err != nil {
		return nil, err
	}

	return &Access{
		AccessToken: Token{
			Token: accessToken,
			TTL:   a.settings.AccessTokenTTL,
		},
		RefreshToken: Token{
			Token: refreshToken,
			TTL:   a.settings.RefreshTokenTTL,
		},
	}, nil
}
