package dbUtils

const (
	NilRows         = Error("Received rows is nil")
	ConnConfigEmpty = Error("connConfig is nil")
)

type Error string

func (e Error) Error() string { return string(e) }
