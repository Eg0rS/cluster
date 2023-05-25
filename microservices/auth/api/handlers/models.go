package handlers

type TokenRequest struct {
	GrantType    GrantType `json:"grant_type"`
	UserName     string    `json:"email"`
	Password     string    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

const (
	PasswordGrantType     = GrantType("password")
	RefreshTokenGrantType = GrantType("refresh_token")
)

type GrantType string
