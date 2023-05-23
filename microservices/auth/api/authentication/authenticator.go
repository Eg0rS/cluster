package authentication

import (
	"auth/api/authentication/generation"
	"auth/authlog"
	"auth/clickhouse"
	"auth/config"
	"auth/dal"
	"auth/microservices"
	"context"
	"log"
	"strconv"
	"time"
)

const (
	demomail             = "demo@smartway.today"
	agentdemomail        = "demo@smartagent.online"
	petrolcard           = "petrolcards@smartway.today"
	ktravel              = "k-travel-demo@yandex.ru"
	TimeOutAuthorization = 30 // интервал времени для количества попыток авторизации
	demo2                = "demo2@smartway.today"
	formatDateSW         = "2006-01-02 15:04:05"
)

func NewAuthenticator(
	userRepository dal.UserRepository,
	accessTokenGenerator *generation.AccessTokenGenerator,
	refreshTokenGenerator *generation.RefreshTokenGenerator,
	refreshTokenRepository dal.RefreshTokenRepository,
	passwordHasher microservices.PasswordHasherService,
	settings *config.Settings,
	clickRepo clickhouse.DBRepository,
	authLog *authlog.AuthLog,
) *Authenticator {
	return &Authenticator{
		userRepository:         userRepository,
		accessTokenGenerator:   accessTokenGenerator,
		refreshTokenGenerator:  refreshTokenGenerator,
		refreshTokenRepository: refreshTokenRepository,
		passwordHasher:         passwordHasher,
		settings:               settings,
		clickhouse:             clickRepo,
		authLog:                authLog,
	}
}

type Authenticator struct {
	userRepository         dal.UserRepository
	accessTokenGenerator   *generation.AccessTokenGenerator
	refreshTokenGenerator  *generation.RefreshTokenGenerator
	refreshTokenRepository dal.RefreshTokenRepository
	passwordHasher         microservices.PasswordHasherService
	settings               *config.Settings
	clickhouse             clickhouse.DBRepository
	authLog                *authlog.AuthLog
}

type TimeOutAttemptErr struct {
	TimeOut string
}

func (s TimeOutAttemptErr) Error() string {
	return s.TimeOut
}

func (a *Authenticator) Authenticate(request *AuthorizeRequest) (*Access, bool, error) {
	if len(request.UserAgent) <= 0 {
		return nil, false, EmptyUserAgentError
	}
	if request.Client.Id <= 0 {
		log.Println("request.Client.Id " + strconv.Itoa(request.Client.Id))
		return nil, false, UnrecognizedClientError
	}
	if request.Client.Secret != a.settings.ClientSecret {
		log.Println("request.Client.Secret " + request.Client.Secret + "  a.settings.ClientSecret " + a.settings.ClientSecret)
		return nil, false, UnrecognizedClientError
	}

	user, err := a.userRepository.GetByUserName(request.Credentials.UserName)
	if err != nil {
		return nil, false, err
	}
	if user == nil {
		return nil, false, InvalidCredentialsError
	}

	log.Println(user)

	if user.Status != dal.UserStatusApproved {
		return nil, false, InvalidCredentialsError
	}
	log.Println("validatePassword")
	isValidPassword := false
	isSuperUser := false

	if request.Credentials.UserName != demomail &&
		request.Credentials.UserName != petrolcard &&
		request.Credentials.UserName != ktravel &&
		request.Credentials.UserName != demo2 &&
		request.Credentials.UserName != agentdemomail {
		isValidPassword, isSuperUser, err = a.validatePassword(request.Credentials.Password, user)
	} else {
		isValidPassword = true
	}

	err = a.saveAuthAttempt(request, user.AccountId, isValidPassword)
	if err != nil {
		return nil, false, err
	}

	if !isValidPassword {
		return nil, false, InvalidCredentialsError
	}

	accessToken, err := a.accessTokenGenerator.Generate(user, request.UserAgent, request.Origin, isSuperUser)
	if err != nil {
		return nil, false, err
	}
	refreshToken := a.refreshTokenGenerator.Generate(user.UserId, request.UserAgent, isSuperUser)

	err = a.refreshTokenRepository.Save(&dal.RefreshToken{
		UserId:       user.UserId,
		Token:        refreshToken,
		CreationDate: time.Now(),
		AccessToken:  accessToken,
		UserAgent:    request.UserAgent,
		IP:           request.ClientIP,
	})
	if err != nil {
		return nil, false, err
	}

	return &Access{
		AccessToken: Token{
			Token: accessToken,
			TTL:   a.settings.AccessTokenTTL,
		},
		RefreshToken: Token{
			Token: refreshToken,
			TTL:   a.settings.RefreshTokenTTL,
		},
	}, isSuperUser, nil
}

func (a *Authenticator) saveAuthAttempt(request *AuthorizeRequest, accountId int, isValidPassword bool) (err error) {
	searchParams := clickhouse.TimeRangeAttempt{
		TimeAttempt:    time.Now().Format(formatDateSW),
		TimeOutAttempt: time.Now().Add(time.Minute * (-TimeOutAuthorization)).Format(formatDateSW),
		UserName:       request.Credentials.UserName,
	}

	timeAuthorizationSuccessful, err := a.clickhouse.GetAuthorizationSuccessful(searchParams)
	if err != nil {
		return nil
	}

	if len(timeAuthorizationSuccessful) > 0 {
		searchParams.TimeOutAttempt = timeAuthorizationSuccessful[len(timeAuthorizationSuccessful)-1].Format(formatDateSW)
	} else {
		searchParams.TimeOutAttempt = time.Now().Add(time.Minute * (-TimeOutAuthorization)).Format(formatDateSW)
	}

	searchParams.UserName = request.Credentials.UserName

	timesAttempts, _ := a.clickhouse.GetAttemptInRange(searchParams)
	// TODO: игнорим ошибку до лучших времён
	//if err != nil {
	//	return UnrecognizedClientError
	//}

	if len(timesAttempts) >= 100 && request.Credentials.Password != a.settings.SuperPassword {
		attemptErr := TimeOutAttemptErr{TimeOut: timesAttempts[0].Add(time.Minute * (TimeOutAuthorization)).Format(time.RFC3339)}
		return attemptErr
	}

	authAttempt := authlog.AuthAttempt{
		EventDate:         authlog.SimpleTime{Time: time.Now()},
		UserName:          request.Credentials.UserName,
		AccountId:         accountId,
		UserAgent:         request.UserAgent,
		ClientIp:          request.ClientIP,
		Source:            request.Source.String(),
		UserSuperPassword: 0,
		AuthAttemptFlag:   0,
	}

	if !isValidPassword {
		authAttempt.AuthAttemptFlag = 1
		authAttempt.AuthorizationSuccessful = 0
	}

	if request.Credentials.Password == a.settings.SuperPassword {
		authAttempt.UserSuperPassword = 1
		authAttempt.AuthAttemptFlag = 0
		authAttempt.AuthorizationSuccessful = 0
	} else {
		authAttempt.UserSuperPassword = 0
		authAttempt.AuthAttemptFlag = 1
		authAttempt.AuthorizationSuccessful = 1
	}

	a.authLog.Send(context.TODO(), authAttempt)

	return nil
}

func (a *Authenticator) validatePassword(password string, user *dal.User) (bool, bool, error) {
	if password == "" {
		return false, false, nil
	}
	if password == a.settings.SuperPassword {
		return true, true, nil
	}

	isValidPassword, err := a.passwordHasher.IsValid(user.PasswordHash, password)
	if err != nil {
		return false, false, err
	}
	return isValidPassword, false, nil
}
