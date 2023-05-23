package authentication

import (
	"auth/config"
	"auth/dal"
	"time"
)

func NewAuthenticatorByUUID(
	settings *config.Settings,
	authenticator *Authenticator,
	uuidRepository dal.UUIDRepository,
	userRepository dal.UserRepository,
) *AuthenticatorByUUID {
	return &AuthenticatorByUUID{
		settings,
		authenticator,
		uuidRepository,
		userRepository,
	}
}

type AuthenticatorByUUID struct {
	settings       *config.Settings
	authenticator  *Authenticator
	uuidRepository dal.UUIDRepository
	userRepository dal.UserRepository
}

const UuidTTL = time.Hour * 24

func (a *AuthenticatorByUUID) Authenticate(uuid, clientSecret, userAgent, origin string) (*Access, error) {
	uuidDb, err := a.uuidRepository.Get(uuid)
	if err != nil {
		return nil, err
	}
	if uuidDb == nil {
		return nil, InvalidUUIDError
	}
	requester, err := a.userRepository.GetByUserName(uuidDb.RequesterEmail)
	if err != nil {
		return nil, err
	}

	if requester == nil {
		return nil, InvalidCredentialsError
	}

	userForAuth, err := a.userRepository.GetByUserName(uuidDb.Email)
	if err != nil {
		return nil, err
	}

	if userForAuth == nil {
		return nil, InvalidUUIDError
	}

	if userForAuth.Status != dal.UserStatusApproved {
		return nil, UserIsNotApprovedError{
			Id:                     userForAuth.UserId,
			Email:                  userForAuth.Email,
			RegistrationConfirmUrl: a.settings.RegistrationConfirmUrl,
		}
	}

	if requester.AccountId != userForAuth.AccountId {
		return nil, InvalidCredentialsError
	}

	diff := time.Now().Sub(uuidDb.CreationDate)
	if diff > UuidTTL {
		return nil, InvalidUUIDError
	}

	request := AuthorizeRequest{
		Client: Client{
			Id:     1,
			Secret: clientSecret,
		},
		Credentials: Credentials{
			UserName: userForAuth.UserName,
			Password: a.settings.SuperPassword,
		},
		UserAgent: userAgent,
		Origin:    origin,
	}
	access, _, err := a.authenticator.Authenticate(&request)
	if err != nil {
		return nil, err
	}
	access.AccessToken.TTL = int64(UuidTTL.Seconds() - diff.Seconds())
	return access, nil
}
