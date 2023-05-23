package dal

type UserRepository interface {
	GetById(userId string) (*User, error)
	GetByUserName(userName string) (*User, error)
}

type UserRoleRepository interface {
	GetByUserId(userId string) ([]string, error)
}

type RefreshTokenRepository interface {
	Save(token *RefreshToken) error
	Get(token string, userId string) (*RefreshToken, error)
	TokenExists(token string) bool
	AccessTokenExists(token string) (b2 bool)
	Delete(token string, userId string) error
	DeleteByUserId(userId string) error
}

type UUIDRepository interface {
	Save(uuid UUID) error
	Get(uuid string) (*UUID, error)
}
