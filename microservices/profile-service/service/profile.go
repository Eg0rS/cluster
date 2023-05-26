package service

import (
	"context"
	"go.uber.org/zap"
	profile "profile_service/client/Profile"
	"profile_service/database/profile_repo"
	"profile_service/model"
)

type ProfileRepo interface {
	UpsertUserInfo(ctx context.Context, model model.UpsertUserInfoModel, userId string) error
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

func (s ProfileService) UpsertUserInfo(ctx context.Context, req profile.UpsertUserInfoReq, userId string) error {
	var data = profile.MapClientToServiceUpsertInfo(req)
	err := s.repo.UpsertUserInfo(ctx, data, userId)
	if err != nil {
		return err
	}

	return nil
}
