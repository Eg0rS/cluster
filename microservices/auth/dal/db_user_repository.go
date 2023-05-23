package dal

import (
	"auth/utils/dbUtils"
	"database/sql"
	"fmt"
)

func NewDbUserRepository(db *sql.DB) UserRepository {
	return &DbUserRepository{db}
}

type DbUserRepository struct {
	db *sql.DB
}

func (r *DbUserRepository) GetById(userId string) (*User, error) {
	user := &User{}
	rows, err := r.db.Query(`
			select top 1
					u.Id,
					u.UserName,
					u.PasswordHash,
					u.Email,
					u.MultiCompanyId,
					u.RegistrationStatus
				from dbo.AspNetUsers as u
				where u.Id = @USER_ID
		`,
		sql.Named("USER_ID", userId),
	)
	if err != nil {
		return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
	}
	if rows == nil {
		return nil, dbUtils.NilRows
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.UserId, &user.UserName, &user.PasswordHash, &user.Email, &user.AccountId, &user.Status)
		if err != nil {
			return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
		}
		return user, nil
	}
	return nil, fmt.Errorf("user not found: %s", userId)
}

func (r *DbUserRepository) GetByUserName(userName string) (*User, error) {
	user := &User{}
	rows, err := r.db.Query(`
			select top 1
					u.Id,
					u.UserName,
					u.PasswordHash,
					u.Email,
					u.MultiCompanyId,
					u.RegistrationStatus
				from dbo.AspNetUsers as u
				where u.UserName = @USER_NAME
		`,
		sql.Named("USER_NAME", userName),
	)
	if err != nil {
		return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
	}
	if rows == nil {
		return nil, dbUtils.NilRows
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.UserId, &user.UserName, &user.PasswordHash, &user.Email, &user.AccountId, &user.Status)
		if err != nil {
			return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
		}

		return user, nil
	}
	return nil, nil
}
