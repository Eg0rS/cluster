package handlers

type TokenRequest struct {
	GrantType       GrantType `json:"grant_type"`
	ClientId        int       `json:"client_id"`
	ClientSecret    string    `json:"client_secret"`
	UserName        string    `json:"username"`
	Password        string    `json:"password"`
	RefreshToken    string    `json:"refresh_token"`
	Source          Source    `json:"source"`
	Uuid            string    `json:"uuid"`
	ClientIP        string
	ClientUserAgent string
	Origin          string
}

type Source string

const (
	SourceMobile   Source = "mobile-api"
	SourceExternal Source = "external-api"
	SourceOther    Source = ""
)

func (s Source) IsMobile() bool {
	return s == SourceMobile
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

const (
	PasswordGrantType     = GrantType("password")
	PasswordHashGrantType = GrantType("password_hash") // TODO It's temporary grant type
	RefreshTokenGrantType = GrantType("refresh_token")
	UuidGrantType         = GrantType("uuid")
)

type GrantType string
