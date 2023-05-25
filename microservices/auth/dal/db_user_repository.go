package dal

import (
	"auth/utils/dbUtils"
	"database/sql"
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
				where u.id = @USER_ID
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
			select top 1
					 * from users as u
				where u.email = @USER_NAME
		`,
		sql.Named("USER_NAME", email),
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
	_, err := r.db.Exec(`
			insert into users
				(email, password_hash, event_date, surname, name, patronymic, city, university, age, education, direction)
			values
				(@EMAIL, @PASSWORD_HASH, @EVENT_DATE, @SURNAME, @NAME, @PATRONYMIC, @CITY, @UNIVERSITY, @AGE, @EDUCATION, @DIRECTION)
		`,
		sql.Named("EMAIL", user.Email),
		sql.Named("PASSWORD_HASH", user.PasswordHash),
		sql.Named("EVENT_DATE", user.EventDate),
		sql.Named("SURNAME", user.Surname),
		sql.Named("NAME", user.Name),
		sql.Named("PATRONYMIC", user.Patronymic),
		sql.Named("CITY", user.City),
		sql.Named("UNIVERSITY", user.University),
		sql.Named("AGE", user.Age),
		sql.Named("EDUCATION", user.Education),
		sql.Named("DIRECTION", user.Direction),
	)
	if err != nil {
		return fmt.Errorf("failed insert into dbo.AspNetUsers: %s", err.Error())
	}
	return nil
}
