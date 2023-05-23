package microservices

import (
	"context"
	"gitea.gospodaprogrammisty.ru/Go/servicelib/logging/http"
)

type PasswordHasherService interface {
	Get(password string) (string, error)
	IsValid(hashedPassword string, password string) (bool, error)
}

type LoggerSlackService interface {
	Send(sendMessage *LoggerSlackTemplateDTO) error
}

type LoggerService interface {
	LogAccess(ctx context.Context, entry http.Entry) error
}
