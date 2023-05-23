package generation

import (
	"auth/cipher"
	"auth/config"
	"encoding/json"
)

func NewRefreshTokenParser(settings *config.Settings) *RefreshTokenParser {
	return &RefreshTokenParser{settings}
}

type RefreshTokenParser struct {
	settings *config.Settings
}

func (r *RefreshTokenParser) Parse(refreshToken string) (*RefreshTokenClaims, error) {
	decodedToken, err := cipher.Decrypt([]byte(r.settings.RefreshTokenSecret), refreshToken)
	if err != nil {
		return nil, err
	}
	var claims RefreshTokenClaims
	if err := json.Unmarshal([]byte(decodedToken), &claims); err != nil {
		return nil, err
	}

	return &claims, nil
}
