package dal

const (
	UserNotFoundError    = Error("User not found")
	RefreshTokenNotFound = Error("Refresh token not found")
)

type Error string

func (e Error) Error() string { return string(e) }
