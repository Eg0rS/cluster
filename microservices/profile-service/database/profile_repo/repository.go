package profile_repo

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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
