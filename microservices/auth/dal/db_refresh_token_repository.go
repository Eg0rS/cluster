package dal

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewDbRefreshTokenRepository(db *sqlx.DB) RefreshTokenRepository {
	return &DbRefreshTokenRepository{db}
}

type DbRefreshTokenRepository struct {
	db *sqlx.DB
}

func (r *DbRefreshTokenRepository) Save(token *RefreshToken) error {
	var resErr error
	_, resErr = r.db.Query(`
			INSERT INTO refresh_tokens (user_id, access_token, refresh_token)
			VALUES ($1, $2, $3)`,
		token.UserId,
		token.AccessToken,
		token.RefreshToken,
	)
	return resErr
}

func (r *DbRefreshTokenRepository) Get(token string, userId int) (*RefreshToken, error) {
	rows, _ := r.db.Query(`
			select id, user_id, access_token, refresh_token, event_date from refresh_tokens as rt
				where rt.refresh_token = $1 and rt.user_id = $2
		`,
		token,
		userId,
	)
	result := RefreshToken{}
	if rows.Next() {
		err := rows.Scan(&result.Id, &result.UserId, &result.AccessToken, &result.RefreshToken, &result.EventDate)
		if err != nil {
			return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
		}
	}
	defer rows.Close()

	return &result, nil
}

func (r *DbRefreshTokenRepository) TokenExists(token string) (b2 bool) {
	rows, _ := r.db.Query(`
			select * from refresh_tokens as rt
				where rt.refresh_token = $1
		`,
		token,
	)
	if rows != nil {
		return true
	}
	defer rows.Close()
	return false
}

func (r *DbRefreshTokenRepository) AccessTokenExists(token string) (b2 bool) {
	rows, _ := r.db.Query(`
			select * from refresh_tokens as rt
				where rt.access_token = $1
		`,
		token,
	)
	if rows != nil {
		return true
	}
	defer rows.Close()
	return false
}

func (r *DbRefreshTokenRepository) DeleteByUserId(userId string) (err error) {
	_, err = r.db.Query(`
			Delete from refresh_tokens as rt
				where rt.user_id =$1
		`,
		userId,
	)

	return err
}
