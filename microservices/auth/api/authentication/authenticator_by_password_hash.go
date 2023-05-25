package authentication

import (
	"auth/api/authentication/generation"
	"auth/config"
	"auth/dal"
	"auth/microservices"
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

	user, err := a.userRepository.GetByUserName(request.Credentials.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, InvalidCredentialsError
	}

	if request.Credentials.Password != user.PasswordHash {
		return nil, InvalidCredentialsError
	}

	isSuperUser := request.Credentials.Password == a.settings.SuperPassword

	accessToken, err := a.accessTokenGenerator.Generate(user, isSuperUser)
	if err != nil {
		return nil, err
	}
	refreshToken := a.refreshTokenGenerator.Generate(user.Id, isSuperUser)

	err = a.refreshTokenRepository.Save(&dal.RefreshToken{
		UserId:       user.Id,
		RefreshToken: refreshToken,
		EventDate:    time.Now(),
		AccessToken:  accessToken,
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
