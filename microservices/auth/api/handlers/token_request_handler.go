package handlers

import (
	"auth/api/authentication"
	"auth/dal"
	"auth/microservices"
	"auth/microservices/loggerSlack"
	"auth/microservices/userservice"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func NewTokenRequestHandler(
	authenticator *authentication.Authenticator,
	authenticatorByPasswordHash *authentication.AuthenticatorByPasswordHash,
	refresher *authentication.Refresher,
	authenticatorByUUID *authentication.AuthenticatorByUUID,
	loggerSlack microservices.LoggerSlackService,
	dbAvailableUserRepo *dal.DbAvailableUserRepository,
) *TokenRequestHandler {
	return &TokenRequestHandler{
		authenticator:               authenticator,
		authenticatorByPasswordHash: authenticatorByPasswordHash,
		refresher:                   refresher,
		authenticatorByUUID:         authenticatorByUUID,
		loggerSlack:                 loggerSlack,
		dbAvailableUserRepo:         dbAvailableUserRepo,
	}
}

type TokenRequestHandler struct {
	authenticator               *authentication.Authenticator
	authenticatorByPasswordHash *authentication.AuthenticatorByPasswordHash
	refresher                   *authentication.Refresher
	authenticatorByUUID         *authentication.AuthenticatorByUUID
	loggerSlack                 microservices.LoggerSlackService
	dbAvailableUserRepo         *dal.DbAvailableUserRepository
}

func (h *TokenRequestHandler) HandleTokenRequest(request *http.Request) (interface{}, error) {
	requestData, err := h.parseRequestData(request)
	if err != nil {
		return nil, err
	}
	switch requestData.GrantType {
	case PasswordGrantType:
		log.Println("Authorize by " + PasswordGrantType)
		tokenResponse, err := h.authByCredentials(requestData)
		if err != nil {
			return nil, err
		}
		return tokenResponse, nil
	case RefreshTokenGrantType:
		log.Println("Authorize by " + RefreshTokenGrantType)
		tokenResponse, err := h.authByRefreshToken(requestData)
		if err != nil {
			return nil, err
		}
		return tokenResponse, nil
	case PasswordHashGrantType: // TODO temp case
		tokenResponse, err := h.authByUserNameAndPasswordHash(requestData)
		if err != nil {
			return nil, err
		}
		return tokenResponse, nil
	case UuidGrantType:
		tokenResponse, err := h.authByUuid(requestData)
		if err != nil {
			return nil, err
		}
		return tokenResponse, nil
	default:
		return nil, BadRequestError
	}
}

func (h *TokenRequestHandler) parseRequestData(request *http.Request) (*TokenRequest, error) {
	if request.Method != http.MethodPost {
		return nil, BadRequestError
	}
	decoder := json.NewDecoder(request.Body)
	var requestData TokenRequest
	err := decoder.Decode(&requestData)
	if err != nil {
		return nil, BadRequestError
	}
	requestData.ClientIP = request.Header.Get("X-Real-Ip")
	requestData.ClientUserAgent = request.Header.Get("User-Agent")
	requestData.Origin = request.Header.Get("Origin")
	defer request.Body.Close()
	return &requestData, nil
}

var (
	alias = map[string]string{
		"1234": "demo@smartway.today",
		"2016": "demo@smartway.today",
		"1C":   "1ctest@smartway.today",
		"1c":   "1ctest@smartway.today",
		"2020": "petrolcards@smartway.today",
		"2021": "k-travel-demo@yandex.ru",
		"2345": "demo2@smartway.today",
	}
)

func (h *TokenRequestHandler) authByCredentials(request *TokenRequest) (*TokenResponse, error) {
	authorizeRequest := authentication.AuthorizeRequest{
		Client: authentication.Client{
			Id:     request.ClientId,
			Secret: request.ClientSecret,
		},
		Credentials: authentication.Credentials{
			UserName: request.UserName,
			Password: request.Password,
		},
		UserAgent: request.ClientUserAgent,
		ClientIP:  request.ClientIP,
		Source:    authentication.Source(request.Source),
		Origin:    request.Origin,
	}

	if userName, ok := alias[authorizeRequest.Credentials.UserName]; ok {
		authorizeRequest.Credentials.UserName = userName
	}

	access, isSuperUser, err := h.authenticator.Authenticate(&authorizeRequest)
	if err != nil {
		log.Println(err)
	}
	log.Println("Authenticate proshla uspeshno")
	if err != nil {
		switch err {
		case authentication.EmptyUserAgentError:
			return nil, BadRequestError
		case authentication.UnrecognizedClientError:
			h.sendMessageInSlack(loggerSlack.UnrecognizedClientMessage, request.UserName, request.Source.IsMobile())
			return nil, ForbiddenError
		case authentication.InvalidClientSecretError:
			return nil, ForbiddenError
		case authentication.InvalidCredentialsError:
			h.sendMessageInSlack(loggerSlack.FailMessage, request.UserName, request.Source.IsMobile())
			return nil, ForbiddenError
		default:
			return nil, err
		}
	}

	messageType := loggerSlack.EnterMessage

	if isSuperUser {
		messageType = loggerSlack.EnterWithSuperPasswordMessage
	}
	h.sendMessageInSlack(messageType, request.UserName, request.Source.IsMobile())
	return &TokenResponse{
		AccessToken:  access.AccessToken.Token,
		ExpiresIn:    access.AccessToken.TTL,
		RefreshToken: access.RefreshToken.Token,
		TokenType:    "bearer",
	}, nil
}

func (h *TokenRequestHandler) authByRefreshToken(request *TokenRequest) (*TokenResponse, error) {
	refreshRequest := authentication.RefreshRequest{
		Client: authentication.Client{
			Id:     request.ClientId,
			Secret: request.ClientSecret,
		},
		RefreshToken: request.RefreshToken,
		UserAgent:    request.ClientUserAgent,
		Origin:       request.Origin,
		ClientIP:     request.ClientIP,
	}
	access, err := h.refresher.RefreshAccess(&refreshRequest)
	if err != nil {
		log.Println(err)
		switch err {
		case authentication.EmptyUserAgentError:
			return nil, BadRequestError
		case authentication.UnrecognizedClientError:
			return nil, ForbiddenError
		case authentication.InvalidClientSecretError:
			return nil, ForbiddenError
		case authentication.InvalidTokenError:
			return nil, ForbiddenError
		default:
			return nil, err
		}
	}

	return &TokenResponse{
		AccessToken:  access.AccessToken.Token,
		ExpiresIn:    access.AccessToken.TTL,
		RefreshToken: access.RefreshToken.Token,
		TokenType:    "bearer",
	}, nil
}

func (h *TokenRequestHandler) authByUuid(request *TokenRequest) (*TokenResponse, error) {
	access, err := h.authenticatorByUUID.Authenticate(request.Uuid, request.ClientSecret, request.ClientUserAgent, request.Origin)
	if err != nil {
		switch err {
		case authentication.EmptyUserAgentError:
			return nil, BadRequestError
		case authentication.UnrecognizedClientError:
			h.sendMessageInSlack(loggerSlack.UnrecognizedClientMessage, request.Uuid, request.Source.IsMobile())
			return nil, ForbiddenError
		case authentication.InvalidClientSecretError:
			return nil, ForbiddenError
		case authentication.InvalidCredentialsError:
			h.sendMessageInSlack(loggerSlack.FailMessage, request.Uuid, request.Source.IsMobile())
			return nil, ForbiddenError
		default:
			return nil, err
		}
	}

	h.sendMessageInSlack(loggerSlack.EnterByUuid, request.Uuid, request.Source.IsMobile())
	return &TokenResponse{
		AccessToken:  access.AccessToken.Token,
		ExpiresIn:    access.AccessToken.TTL,
		RefreshToken: access.RefreshToken.Token,
		TokenType:    "bearer",
	}, nil
}

// TODO temp method
func (h *TokenRequestHandler) authByUserNameAndPasswordHash(request *TokenRequest) (*TokenResponse, error) {
	authorizeRequest := authentication.AuthorizeRequest{
		Client: authentication.Client{
			Id:     request.ClientId,
			Secret: request.ClientSecret,
		},
		Credentials: authentication.Credentials{
			UserName: request.UserName,
			Password: request.Password,
		},
		UserAgent: request.ClientUserAgent,
		Origin:    request.Origin,
	}

	access, err := h.authenticatorByPasswordHash.Authenticate(&authorizeRequest)
	if err != nil {
		switch err {
		case authentication.EmptyUserAgentError:
			return nil, BadRequestError
		case authentication.UnrecognizedClientError:
			return nil, ForbiddenError
		case authentication.InvalidClientSecretError:
			return nil, ForbiddenError
		case authentication.InvalidCredentialsError:
			return nil, ForbiddenError
		default:
			return nil, err
		}
	}

	return &TokenResponse{
		AccessToken:  access.AccessToken.Token,
		ExpiresIn:    access.AccessToken.TTL,
		RefreshToken: access.RefreshToken.Token,
		TokenType:    "bearer",
	}, nil
}

func (h *TokenRequestHandler) sendMessageInSlack(messageType loggerSlack.SlackMessageType, userName string, mobile bool) {
	msg := ""
	msgType := "SECURITY"

	log.Println("Message type: ", loggerSlack.EnterMessage)

	application := "систему"
	if mobile {
		application = "в мобильное приложение"
	}

	switch messageType {
	case loggerSlack.EnterMessage:
		msg = fmt.Sprintf("Пользователь %s выполнил вход в %s", userName, application)
		break
	case loggerSlack.EnterWithSuperPasswordMessage:
		msgType = "SUPERUSER"
		msg = fmt.Sprintf("Пользователь %s выполнил вход в %s с суперпаролем", userName, application)
		break
	case loggerSlack.UnrecognizedClientMessage:
		msg = fmt.Sprintf("Попытка входа не зарегестрированного пользователя %s", userName)
		break
	case loggerSlack.FailMessage:
		msg = fmt.Sprintf("Неудачная попытка входа в %s. Логин => %s", application, userName)
		break
	}

	if len(msg) != 0 {
		msg := &microservices.LoggerSlackTemplateDTO{
			Text: msg,
			Type: msgType,
		}

		loggerSlackErr := h.loggerSlack.Send(msg)

		if loggerSlackErr != nil {
			fmt.Println(loggerSlackErr.Error())
		}
	}
}

func (h *TokenRequestHandler) MobileUserValidation(request *TokenRequest) (err error) {
	if !request.Source.IsMobile() {
		return nil
	}

	var errSuffix string
	defer func() {
		if err != nil {
			_ = h.loggerSlack.Send(&microservices.LoggerSlackTemplateDTO{
				Type: "MOBILE_LOGIN_RESTRICT",
				Text: fmt.Sprintf("Попытка входа в мобильное приложение неодобренного пользователя %s.\n%s", request.UserName, errSuffix),
			})

			err := h.dbAvailableUserRepo.SaveUser(request.UserName, err.Error())
			if err != nil {
				log.Println(err)
			}
		}
	}()

	if strings.HasSuffix(request.UserName, "@smartway.today") ||
		h.dbAvailableUserRepo.UserIsAvailable(request.UserName) {
		return nil
	}

	info, err := userservice.GetUserInfo(request.UserName)
	if err != nil {
		errSuffix = "не найдены права для " + request.UserName
		return err
	}

	// Если пользователь не связан с сотрудником, то разрешаем вход
	if info.EmployeeLink == 0 {
		return nil
	}

	// Если связан и не нашли прав, запрещаем вход
	if info.Rights == nil {
		errSuffix = fmt.Sprintf("Не найдены права.")
		return fmt.Errorf("неправильные права у пользователя [%s]", request.UserName)
	}

	log.Printf("%s = %d", request.UserName, info.Rights.BuyTripAccount)
	// разрешаем вход если есть право бронировать для всех
	if info.Rights.BuyTripAccount == 3 {
		return nil
	} else {
		errSuffix = fmt.Sprintf("Нет прав бронировать всем")
		return errors.New("запрещён вход для мобильного приложения")
	}

	return errors.New("запрещён вход для мобильного приложения")
}
