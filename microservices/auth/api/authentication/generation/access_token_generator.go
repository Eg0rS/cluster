package generation

import (
	"auth/config"
	"auth/dal"
	"auth/jwt"
	"auth/microservices/accountservice"
	"errors"
	"log"
	"regexp"
	"time"
)

func NewAccessTokenGenerator(
	settings *config.Settings,
	roleChecker *RoleChecker,
	accountService accountservice.Service,
) *AccessTokenGenerator {
	return &AccessTokenGenerator{settings, roleChecker, accountService}
}

type AccessTokenGenerator struct {
	settings       *config.Settings
	roleChecker    *RoleChecker
	accountService accountservice.Service
}

func (g *AccessTokenGenerator) Generate(user *dal.User, userAgent string, origin string, isSuperUser bool) (string, error) {
	isBackofficeManager, err := g.roleChecker.IsBackofficeManager(user.UserId)
	if err != nil {
		return "", err
	}

	isSmartAgent, err := g.accountService.IsSmartAgent(user.AccountId)
	if err != nil {
		log.Println("не удалось проверить аккаунт на смартагент")
		return "", err
	}

	isSmartAgentOrigin := IsSmartAgentOrigin(origin)
	if !isSmartAgent && isSmartAgentOrigin {
		return "", errors.New("запрещен вход в смартагент пользователю смартвей")
	}

	if isSmartAgent && !isBackofficeManager && !isSmartAgentOrigin {
		return "", errors.New("запрещен вход в смартвей пользователю смартагента")
	}

	accessTokenClaims := AccessTokenClaims{
		UserId:              user.UserId,
		UserName:            user.UserName,
		Email:               user.Email,
		AccountId:           user.AccountId,
		IsBackofficeManager: isBackofficeManager,
		UserAgent:           userAgent,
		CreationTimestamp:   time.Now().UTC().Unix(),
		TTL:                 g.settings.AccessTokenTTL,
		US:                  isSuperUser,
	}
	return jwt.GetToken(accessTokenClaims, g.settings.JwtSecret), nil
}

func IsSmartAgentOrigin(origin string) bool {
	for _, pattern := range []string{`https?://.*?\.?smartagent\.online`, `https?://.*?\.sadev`} {
		if match, _ := regexp.MatchString(pattern, origin); match {
			return true
		}
	}

	return false
}
