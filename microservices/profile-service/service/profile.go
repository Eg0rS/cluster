package service

import (
	"go.uber.org/zap"
	"profile_service/database/profile_repo"
)

type ProfileRepo interface {
}

type ProfileService struct {
	logger *zap.SugaredLogger
	repo   ProfileRepo
}

func NewProfileService(logger *zap.SugaredLogger, repo profile_repo.ProfileRepository) ProfileService {
	return ProfileService{
		logger: logger,
		repo:   repo,
	}
}
