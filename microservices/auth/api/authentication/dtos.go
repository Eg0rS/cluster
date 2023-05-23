package authentication

type AuthorizeRequest struct {
	Client      Client
	Credentials Credentials
	UserAgent   string
	ClientIP    string
	Source      Source
	Origin      string
}

type Source string

func (s Source) String() string {
	if s == SourceOther {
		return "other"
	}

	return string(s)
}

const (
	SourceMobile   Source = "mobile-api"
	SourceExternal Source = "external-api"
	SourceOther    Source = ""
)

func (s Source) IsMobile() bool {
	return s == SourceMobile
}

type RefreshRequest struct {
	Client       Client
	RefreshToken string
	UserAgent    string
	Origin       string
	ClientIP     string
}

type Client struct {
	Id     int
	Secret string
}

type Credentials struct {
	UserName string
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
