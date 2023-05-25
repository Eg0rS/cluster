package dal

type UserRepository interface {
	GetById(userId int) (*User, error)
	GetByUserName(userName string) (*User, error)
	Create(user *User) error
}

type UserRoleRepository interface {
	GetByUserId(userId string) ([]string, error)
}

type RefreshTokenRepository interface {
	Save(token *RefreshToken) error
	Get(token string, userId int) (*RefreshToken, error)
	TokenExists(token string) bool
	AccessTokenExists(token string) (b2 bool)
	DeleteByUserId(userId string) error
}
