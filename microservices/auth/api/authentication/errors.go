package authentication

const (
	UnrecognizedClientError  = Error("Unrecognized client")
	InvalidClientSecretError = Error("Invalid client secret")
	InvalidCredentialsError  = Error("Invalid credentials")
	EmptyUserAgentError      = Error("Empty user agent")
	InvalidTokenError        = Error("Invalid token")
	InvalidUUIDError         = Error("Invalid uuid")
)

type Error string

func (e Error) Error() string { return string(e) }

type UserIsNotApprovedError struct {
	Email                  string
	Id                     string
	RegistrationConfirmUrl string
}

func (u UserIsNotApprovedError) Error() string {
	return "User is not approved"
}
