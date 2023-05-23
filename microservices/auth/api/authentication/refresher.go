package authentication

import (
	"auth/api/authentication/generation"
	"auth/config"
	"auth/dal"
	"log"
	"time"
)

func NewRefresher(
	repository dal.UserRepository,
	accessTokenGenerator *generation.AccessTokenGenerator,
	refreshTokenGenerator *generation.RefreshTokenGenerator,
	tokenParser *generation.RefreshTokenParser,
	refreshTokenRepository dal.RefreshTokenRepository,
	settings *config.Settings,
) *Refresher {
	return &Refresher{
		repository,
		accessTokenGenerator,
		refreshTokenGenerator,
		tokenParser,
		refreshTokenRepository,
		settings,
	}
}

type Refresher struct {
	repository             dal.UserRepository
	accessTokenGenerator   *generation.AccessTokenGenerator
	refreshTokenGenerator  *generation.RefreshTokenGenerator
	tokenParser            *generation.RefreshTokenParser
	refreshTokenRepository dal.RefreshTokenRepository
	settings               *config.Settings
}

func (r *Refresher) RefreshAccess(request *RefreshRequest) (*Access, error) {
	if len(request.UserAgent) <= 0 {
		return nil, EmptyUserAgentError
	}

	if request.Client.Id <= 0 {
		return nil, UnrecognizedClientError
	}

	if request.Client.Secret != r.settings.ClientSecret {
		return nil, UnrecognizedClientError
	}

	claims, err := r.tokenParser.Parse(request.RefreshToken)
	if err != nil {
		return nil, InvalidTokenError
	}

	//if claims.UserAgent != request.UserAgent {
	//	log.Println("claims.UserAgent " + claims.UserAgent + "  request.UserAgent " + request.UserAgent)
	//	log.Println(UnrecognizedClientError)
	//	return nil, UnrecognizedClientError
	//}

	if claims.IsExpired() {
		return nil, InvalidTokenError
	}

	_, err = r.refreshTokenRepository.Get(request.RefreshToken, claims.UserId)
	if err != nil {
		log.Println(err)
		return nil, InvalidTokenError
	}

	user, err := r.repository.GetById(claims.UserId)
	if err != nil {
		log.Println(err)
		return nil, InvalidTokenError
	}

	accessToken, err := r.accessTokenGenerator.Generate(user, request.UserAgent, request.Origin, claims.US)
	if err != nil {
		return nil, err
	}
	refreshToken := r.refreshTokenGenerator.Generate(user.UserId, request.UserAgent, claims.US)

	err = r.refreshTokenRepository.Save(&dal.RefreshToken{
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
			TTL:   r.settings.AccessTokenTTL,
		},
		RefreshToken: Token{
			Token: refreshToken,
			TTL:   r.settings.RefreshTokenTTL,
		},
	}, nil
}
