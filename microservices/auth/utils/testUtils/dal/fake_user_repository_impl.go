package dal

import (
	"auth/dal"
)

func NewFakeUserRepository() dal.UserRepository {
	return &FakeUserRepository{}
}

type FakeUserRepository struct{}

func (r *FakeUserRepository) GetById(userId string) (*dal.User, error) {
	return &dal.User{
		UserId:       userId,
		PasswordHash: "HASHED_password",
		AccountId:    42,
		UserName:     "user",
		Status:       dal.UserStatusApproved,
	}, nil
}

func (r *FakeUserRepository) GetByUserName(userName string) (*dal.User, error) {
	if userName == "NOT_APPROVED_USER" {
		return &dal.User{
			UserId:       "123",
			PasswordHash: "HASHED_password",
			AccountId:    42,
			UserName:     "user",
			Status:       dal.UserStatusPending,
		}, nil
	}

	if userName != "user" {
		return nil, nil
	}
	return &dal.User{
		UserId:       "123",
		PasswordHash: "HASHED_password",
		AccountId:    42,
		UserName:     "user",
		Status:       dal.UserStatusApproved,
	}, nil
}
