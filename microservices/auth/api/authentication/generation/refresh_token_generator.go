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

func (g *RefreshTokenGenerator) Generate(userId int, isSuperUser bool) string {
	content := g.getContent(userId, isSuperUser)
	token, err := cipher.Encrypt([]byte(g.settings.RefreshTokenSecret), content)
	if err != nil {
		panic(err) // TODO don't fail
	}
	return token
}

func (g *RefreshTokenGenerator) getContent(userId int, isSuperUser bool) []byte {
	claims := RefreshTokenClaims{
		UserId:            userId,
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
