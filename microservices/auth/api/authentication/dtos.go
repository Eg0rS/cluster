package authentication

type AuthorizeRequest struct {
	Credentials Credentials
	ClientIP    string
}

type RefreshRequest struct {
	RefreshToken string
}

type Credentials struct {
	Email    string
	Password string
}

type Token struct {
	Token string
	TTL   int64
}

type Access struct {
	AccessToken  Token
	RefreshToken Token
}
