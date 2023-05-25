package dal

import "time"

type User struct {
	Id           int
	Email        string
	EventDate    time.Time
	PasswordHash string
	Surname      string
	Name         string
	Patronymic   string
	City         string
	University   string
	Age          int
	Education    string
	Direction    string
}

type RefreshToken struct {
	Id           int
	UserId       int
	RefreshToken string
	EventDate    time.Time
	AccessToken  string
}
