package dal

import (
	"auth/utils/dbUtils"
	"database/sql"
	"fmt"
)

func NewDbUserRoleRepository(db *sql.DB) UserRoleRepository {
	return &DbUserRoleRepository{db}
}

type DbUserRoleRepository struct {
	db *sql.DB
}

func (r *DbUserRoleRepository) GetByUserId(userId string) ([]string, error) {
	result := make([]string, 0)
	rows, err := r.db.Query(`
			select r.Name
				from
					dbo.AspNetUserRoles as ur
					inner join dbo.AspNetRoles as r on r.Id = ur.RoleId
				where ur.UserId = @USER_ID
		`,
		sql.Named("USER_ID", userId),
	)
	if err != nil {
		return nil, fmt.Errorf("failed select from dbo.AspNetRoles: %s", err.Error())
	}
	if rows == nil {
		return nil, dbUtils.NilRows
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, fmt.Errorf("failed select from dbo.AspNetRoles: %s", err.Error())
		}
		result = append(result, name)
	}

	return result, nil
}
