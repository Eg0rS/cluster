package generation

import (
	"auth/config"
	"auth/dal"
	"auth/jwt"
	"time"
)

func NewAccessTokenGenerator(
	settings *config.Settings,

) *AccessTokenGenerator {
	return &AccessTokenGenerator{settings}
}

type AccessTokenGenerator struct {
	settings *config.Settings
}

func (g *AccessTokenGenerator) Generate(user *dal.User, isSuperUser bool) (string, error) {
	accessTokenClaims := AccessTokenClaims{
		UserId:            user.Id,
		Email:             user.Email,
		CreationTimestamp: time.Now().UTC().Unix(),
		TTL:               g.settings.AccessTokenTTL,
		US:                isSuperUser,
	}
	return jwt.GetToken(accessTokenClaims, g.settings.JwtSecret), nil
}
