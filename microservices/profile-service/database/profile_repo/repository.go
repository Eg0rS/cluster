package profile_repo

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"profile_service/database/profile_repo/query"
	"profile_service/model"
)

type ProfileRepository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewProfileRepository(logger *zap.SugaredLogger, db *sqlx.DB) ProfileRepository {
	return ProfileRepository{
		logger: logger,
		db:     db,
	}
}

func (r ProfileRepository) UpsertUserInfo(ctx context.Context, model model.UpsertUserInfoModel, userId string) error {
	if len(model.FirstName) != 0 {
		_, err := r.db.ExecContext(ctx, query.UpsertUserName, model.FirstName, userId)
		if err != nil {
			return err
		}
	}

	if len(model.Surname) != 0 {
		_, err := r.db.ExecContext(ctx, query.UpsertUserSurname, model.Surname, userId)
		if err != nil {
			return err
		}
	}

	if len(model.Patronymic) != 0 {
		_, err := r.db.ExecContext(ctx, query.UpsertUserPatronymic, model.Patronymic, userId)
		if err != nil {
			return err
		}
	}

	if model.Age != 0 {
		_, err := r.db.ExecContext(ctx, query.UpsertUserAge, model.Age, userId)
		if err != nil {
			return err
		}
	}

	if len(model.Education) != 0 {
		_, err := r.db.ExecContext(ctx, query.UpsertUserEducation, model.Education, userId)
		if err != nil {
			return err
		}
	}

	if len(model.University) != 0 {
		_, err := r.db.ExecContext(ctx, query.UpsertUserUniversity, model.University, userId)
		if err != nil {
			return err
		}
	}

	if len(model.City) != 0 {
		_, err := r.db.ExecContext(ctx, query.UpsertUserCity, model.City, userId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r ProfileRepository) SelectUserInfo(ctx context.Context, refreshToken string) (model.UpsertUserInfoModel, error) {
	var userInfo model.UpsertUserInfoModel
	var userId int
	var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	//err := r.db.SelectContext(ctx, &userId, query.GetUserIdByTokenSql, refreshToken)
	//if err != nil {
	//	return model.UpsertUserInfoModel{}, err
	//}

	//SELECT user_id FROM refresh_tokens where refresh_token = ?;

	builder := psql.Select("user_id").From("refresh_tokens").Where("refresh_token", "=", refreshToken).RunWith(r.db)
	err := builder.Scan(&userId)
	if err != nil {
		return model.UpsertUserInfoModel{}, err
	}

	builder = psql.Select("*").From("Organizations").Where("id", "=", userId).RunWith(r.db)
	err = builder.Scan(&userInfo)
	if err != nil {
		return model.UpsertUserInfoModel{}, err
	}
	////err = r.db.Select(ctx, &userInfo, builder, args)
	//if err != nil {
	//	return model.UpsertUserInfoModel{}, err
	//}

	return userInfo, nil
}
