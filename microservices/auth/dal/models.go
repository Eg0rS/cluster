package dal

import "time"

type User struct {
	UserId       string
	UserName     string
	PasswordHash string
	Email        string
	AccountId    int
	Status       UserStatus
}

const (
	UserStatusCreated  = UserStatus(0)
	UserStatusPending  = UserStatus(1)
	UserStatusApproved = UserStatus(2)
	UserStatusRejected = UserStatus(3)
	UserStatusDisabled = UserStatus(4)
)

type UserStatus int

type RefreshToken struct {
	UserId       string    `bson:"user_id"`
	Token        string    `bson:"token"`
	CreationDate time.Time `bson:"creation_date"`
	AccessToken  string    `bson:"access_token"`
	UserAgent    string    `bson:"user_agent"`
	IP           string    `bson:"ip"`
}

type UUID struct {
	UUID           string    `bson:"uuid"`
	Email          string    `bson:"email"`
	RequesterEmail string    `bson:"requester_email"`
	CreationDate   time.Time `bson:"creation_date"`
}
