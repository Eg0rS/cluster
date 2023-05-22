package authorization

import (
	"auth-gateway/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func NewJWTToken(token string, settings *config.Settings) (*JWT, error) {
	if token == "" {
		panic("Empty token")
	}
	jwtParts := strings.Split(token, ".")
	if len(jwtParts) != 3 {
		return nil, fmt.Errorf("Invalid JWT")
	}
	encodedPayloadBytes, err := base64.StdEncoding.DecodeString(jwtParts[1])
	if err != nil {
		return nil, fmt.Errorf("Invalid JWT payload encoding")
	}
	var claims tokenClaims
	if err := json.Unmarshal(encodedPayloadBytes, &claims); err != nil {
		return nil, fmt.Errorf("Invalid JWT payload JSON")
	}
	return &JWT{
		settings:       settings,
		header:         jwtParts[0],
		payload:        jwtParts[1],
		encodedPayload: string(encodedPayloadBytes),
		claims:         claims,
		signature:      jwtParts[2],
	}, nil
}

type JWT struct {
	settings       *config.Settings
	header         string
	payload        string
	claims         tokenClaims
	signature      string
	encodedPayload string
}

func (t *JWT) IsValid() bool {
	if !t.isHeaderValid() {
		return false
	}
	if !t.isSignatureAuthentic() {
		return false
	}
	if t.isExpired() {
		return false
	}
	return true
}

func (t *JWT) IsUserAgentEqual(userAgent string) bool {
	return t.claims.UserAgent == userAgent
}

func (t *JWT) GetData() tokenData {
	return tokenData{
		AccountId: t.claims.AccountId,
		UserId:    t.claims.UserId,
		UserName:  t.claims.UserName,
	}
}

func (t *JWT) GetPayload() string {
	return t.encodedPayload
}

func (t *JWT) isHeaderValid() bool {
	if t.header == "" || t.payload == "" || t.signature == "" {
		return false
	}
	encodedHeader, err := base64.StdEncoding.DecodeString(t.header)
	if err != nil {
		return false
	}
	if string(encodedHeader) != t.settings.JWTHeader {
		return false
	}
	return true
}

func (t *JWT) isSignatureAuthentic() bool {
	computedSignature := computeSignature(t.header, t.payload, t.settings.TokenSecret)
	return computedSignature == t.signature
}

func (t *JWT) isExpired() bool {
	return t.claims.CreationTimestamp+t.claims.TTL < time.Now().UTC().Unix()
}

func computeSignature(header string, payload string, secret string) string {
	secretBase64 := base64.StdEncoding.EncodeToString([]byte(secret))
	signatureContent := fmt.Sprintf("%s.%s", header, payload)
	return computeHmacSha256(signatureContent, secretBase64)
}

func computeHmacSha256(content string, secret string) string {
	hasher := hmac.New(sha256.New, []byte(secret))
	hasher.Write([]byte(content))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

type tokenClaims struct {
	UserAgent         string `json:"user_agent"`
	CreationTimestamp int64  `json:"iat"`
	TTL               int64  `json:"exp"`
	AccountId         int    `json:"account_id"`
	UserId            string `json:"user_id"`
	UserName          string `json:"user_name"`
}

type tokenData struct {
	AccountId int
	UserId    string
	UserName  string
}
