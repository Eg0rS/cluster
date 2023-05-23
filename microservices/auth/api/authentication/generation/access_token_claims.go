package generation

type AccessTokenClaims struct {
	UserId              string `json:"user_id"`
	UserName            string `json:"user_name"`
	Email               string `json:"email"`
	AccountId           int    `json:"account_id"`
	IsBackofficeManager bool   `json:"is_backoffice_manager"`
	UserAgent           string `json:"user_agent"`
	CreationTimestamp   int64  `json:"iat"`
	TTL                 int64  `json:"exp"`
	US                  bool   `json:"us"`
}
