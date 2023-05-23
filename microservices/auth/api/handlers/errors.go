package handlers

const (
	BadRequestError   = Error("Bad request")
	UnauthorizedError = Error("Unauthorized")
	ForbiddenError    = Error("Forbidden")
)

type Error string

func (e Error) Error() string { return string(e) }
