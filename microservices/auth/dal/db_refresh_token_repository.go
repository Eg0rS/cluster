package dal

import (
	"database/sql"
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
			VALUES (@USERID, @ACCESSTOKEN, @REFRSHTOKEN)`,
		sql.Named("USERID", token.UserId),
		sql.Named("ACCESSTOKEN", token.AccessToken),
		sql.Named("REFRSHTOKEN", token.RefreshToken),
	)
	return resErr
}

func (r *DbRefreshTokenRepository) Get(token string, userId int) (result *RefreshToken, err error) {
	rows, _ := r.db.Query(`
			select * from refresh_tokens as rt
				where rt.refresh_token = @TOKEN and rt.user_id = @USERID
		`,
		sql.Named("TOKEN", token),
		sql.Named("USERID", userId),
	)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&result.Id, &result.UserId, &result.AccessToken, &result.RefreshToken, &result.EventDate)
		if err != nil {
			return nil, fmt.Errorf("failed select from dbo.AspNetUsers: %s", err.Error())
		}
	}

	return
}

func (r *DbRefreshTokenRepository) TokenExists(token string) (b2 bool) {
	rows, _ := r.db.Query(`
			select * from refresh_tokens as rt
				where rt.refresh_token = @TOKEN
		`,
		sql.Named("TOKEN", token),
	)
	if rows == nil {
		return true
	}
	defer rows.Close()
	return false
}

func (r *DbRefreshTokenRepository) AccessTokenExists(token string) (b2 bool) {
	rows, _ := r.db.Query(`
			select * from refresh_tokens as rt
				where rt.access_token = @TOKEN
		`,
		sql.Named("TOKEN", token),
	)
	if rows == nil {
		return true
	}
	defer rows.Close()
	return false
}

func (r *DbRefreshTokenRepository) DeleteByUserId(userId string) (err error) {
	_, err = r.db.Query(`
			Delete from refresh_tokens as rt
				where rt.user_id = @USERID
		`,
		sql.Named("USERID", userId),
	)

	return err
}
