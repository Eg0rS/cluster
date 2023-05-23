package generation

import (
	"auth/dal"
)

func NewRoleChecker(
	userRoleRepository dal.UserRoleRepository,
) *RoleChecker {
	return &RoleChecker{userRoleRepository}
}

type RoleChecker struct {
	userRoleRepository dal.UserRoleRepository
}

func (g *RoleChecker) IsBackofficeManager(userId string) (bool, error) {
	roles, err := g.userRoleRepository.GetByUserId(userId)
	if err != nil {
		return false, err
	}

	for _, role := range roles {
		if role == "BackofficeManager" {
			return true, nil
		}
	}

	return false, nil
}
