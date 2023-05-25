package generation

type AccessTokenClaims struct {
	UserId            int    `json:"user_id"`
	Email             string `json:"email"`
	CreationTimestamp int64  `json:"iat"`
	TTL               int64  `json:"exp"`
	US                bool   `json:"us"`
}
