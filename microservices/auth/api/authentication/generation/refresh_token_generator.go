package generation

import (
	"auth/cipher"
	"auth/config"
	"encoding/json"
	"time"
)

func NewRefreshTokenGenerator(settings *config.Settings) *RefreshTokenGenerator {
	return &RefreshTokenGenerator{settings}
}

type RefreshTokenGenerator struct {
	settings *config.Settings
}

func (g *RefreshTokenGenerator) Generate(userId string, userAgent string, isSuperUser bool) string {
	content := g.getContent(userId, userAgent, isSuperUser)
	token, err := cipher.Encrypt([]byte(g.settings.RefreshTokenSecret), content)
	if err != nil {
		panic(err) // TODO don't fail
	}
	return token
}

func (g *RefreshTokenGenerator) getContent(userId string, userAgent string, isSuperUser bool) []byte {
	claims := RefreshTokenClaims{
		UserId:            userId,
		UserAgent:         userAgent,
		TTL:               g.settings.RefreshTokenTTL,
		CreationTimestamp: time.Now().UTC().Unix(),
		US:                isSuperUser,
	}
	contentJson, err := json.Marshal(claims)
	if err != nil {
		panic(err) // TODO don't fail
	}
	return contentJson
}
