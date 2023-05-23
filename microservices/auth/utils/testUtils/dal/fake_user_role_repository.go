package dal

import (
	"auth/dal"
)

func NewFakeUserRoleRepository() dal.UserRoleRepository {
	return &FakeUserRoleRepository{}
}

type FakeUserRoleRepository struct{}

func (r *FakeUserRoleRepository) GetByUserId(userId string) ([]string, error) {
	return []string{"BackofficeManager"}, nil
}
