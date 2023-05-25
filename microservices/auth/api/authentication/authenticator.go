package authentication

import (
	"auth/api/authentication/generation"
	"auth/config"
	"auth/dal"
	"auth/microservices"
	"log"
	"time"
)

func NewAuthenticator(
	userRepository dal.UserRepository,
	accessTokenGenerator *generation.AccessTokenGenerator,
	refreshTokenGenerator *generation.RefreshTokenGenerator,
	refreshTokenRepository dal.RefreshTokenRepository,
	passwordHasher microservices.PasswordHasherService,
	settings *config.Settings,

) *Authenticator {
	return &Authenticator{
		userRepository:         userRepository,
		accessTokenGenerator:   accessTokenGenerator,
		refreshTokenGenerator:  refreshTokenGenerator,
		refreshTokenRepository: refreshTokenRepository,
		passwordHasher:         passwordHasher,
		settings:               settings,
	}
}

type Authenticator struct {
	userRepository         dal.UserRepository
	accessTokenGenerator   *generation.AccessTokenGenerator
	refreshTokenGenerator  *generation.RefreshTokenGenerator
	refreshTokenRepository dal.RefreshTokenRepository
	passwordHasher         microservices.PasswordHasherService
	settings               *config.Settings
}

type TimeOutAttemptErr struct {
	TimeOut string
}

func (s TimeOutAttemptErr) Error() string {
	return s.TimeOut
}

func (a *Authenticator) Authenticate(request *AuthorizeRequest) (*Access, bool, error) {

	user, err := a.userRepository.GetByUserName(request.Credentials.Email)
	if err != nil {
		return nil, false, err
	}
	if user == nil {
		return nil, false, InvalidCredentialsError
	}

	log.Println(user)

	log.Println("validatePassword")
	isValidPassword := false
	isSuperUser := false

	if err != nil {
		return nil, false, err
	}
	isValidPassword, isSuperUser, err = a.validatePassword(request.Credentials.Password, user)
	if !isValidPassword {
		return nil, false, InvalidCredentialsError
	}

	accessToken, err := a.accessTokenGenerator.Generate(user, isSuperUser)
	if err != nil {
		return nil, false, err
	}
	refreshToken := a.refreshTokenGenerator.Generate(user.Id, isSuperUser)

	err = a.refreshTokenRepository.Save(&dal.RefreshToken{
		UserId:       user.Id,
		RefreshToken: refreshToken,
		EventDate:    time.Now(),
		AccessToken:  accessToken,
	})
	if err != nil {
		return nil, false, err
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
	}, isSuperUser, nil
}

func (a *Authenticator) validatePassword(password string, user *dal.User) (bool, bool, error) {
	if password == "" {
		return false, false, nil
	}
	if password == a.settings.SuperPassword {
		return true, true, nil
	}

	isValidPassword := user.PasswordHash == password
	return isValidPassword, false, nil
}
