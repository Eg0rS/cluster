package generation

import "time"

type RefreshTokenClaims struct {
	UserId            string `json:"user_id"`
	UserAgent         string `json:"user_agent"`
	TTL               int64  `json:"ttl"`
	CreationTimestamp int64  `json:"creation_timestamp"`
	US                bool   `json:"us"`
}

func (c *RefreshTokenClaims) IsExpired() bool {
	return (c.CreationTimestamp + c.TTL) < time.Now().UTC().Unix()
}
