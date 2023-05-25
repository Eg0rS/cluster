package dal

import (
	"auth/utils/dbUtils"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewDbUserRepository(db *sqlx.DB) UserRepository {
	return &DbUserRepository{db}
}

type DbUserRepository struct {
	db *sqlx.DB
}

func (r *DbUserRepository) GetById(userId int) (*User, error) {
	user := &User{}
	rows, err := r.db.Query(`
			select top 1 * from users as u
				where u.id = $1
		`,
		userId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
	}
	if rows == nil {
		return nil, dbUtils.NilRows
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.EventDate, &user.Surname,
			&user.Name,
			&user.Patronymic,
			&user.City,
			&user.University,
			&user.Age,
			&user.Education,
			&user.Direction,
		)
		if err != nil {
			return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
		}
		return user, nil
	}
	return nil, fmt.Errorf("user not found: %s", userId)
}

func (r *DbUserRepository) GetByUserName(email string) (*User, error) {
	user := &User{}
	rows, err := r.db.Query(`
			select 
					 * from users as u
				where u.email = $1
		`,
		email,
	)
	if err != nil {
		return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
	}
	if rows == nil {
		return nil, dbUtils.NilRows
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.EventDate, &user.Surname,
			&user.Name,
			&user.Patronymic,
			&user.City,
			&user.University,
			&user.Age,
			&user.Education,
			&user.Direction,
		)
		if err != nil {
			return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
		}
		return user, nil
	}
	return nil, nil
}

func (r *DbUserRepository) Create(user *User) error {
	_, err := r.db.Query(`insert into users (email, password_hash, surname, first_name, patronymic, city, university, age, education, direction_internship)
values
				($1, $2,  $3, $4, $5, $6, $7, $8, $9, $10)`,
		user.Email,
		user.PasswordHash,
		user.Surname,
		user.Name,
		user.Patronymic,
		user.City,
		user.University,
		user.Age,
		user.Education,
		user.Direction,
	)
	if err != nil {
		return fmt.Errorf("failed insert into dbo.AspNetUsers: %s", err.Error())
	}
	return nil
}
