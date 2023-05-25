package authentication

const (
	InvalidClientSecretError = Error("Invalid client secret")
	InvalidCredentialsError  = Error("Invalid credentials")
	InvalidTokenError        = Error("Invalid token")
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
