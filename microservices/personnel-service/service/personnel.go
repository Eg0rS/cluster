package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	organization "personnel_service/client/Organization"
	personnel "personnel_service/client/Personnel"
	"personnel_service/database/pesrsonnel_repo"
	"personnel_service/model"
	"time"
)

type PersonnelRepo interface {
	InsertRadioTest(ctx context.Context, testModel model.RadioTest) (int, error)
	InsertTextTest(ctx context.Context, testModel model.CreateTextTest, filePath string) (int, error)
	InsertRequest(ctx context.Context, reqModel model.Request) error
	GetAllRequests(ctx context.Context, userId string) ([]pesrsonnel_repo.Request, error)
	GetTestById(ctx context.Context, testId string) (model.RadioTest, error)
	InsertOrganization(ctx context.Context, organizationModel model.AddOrganizationModel) (int, error)
	SelectOrganizations(ctx context.Context) (model.GetOrganizationsModel, error)
}

type PersonnelService struct {
	logger *zap.SugaredLogger
	repo   PersonnelRepo
}

func NewPersonnelService(logger *zap.SugaredLogger, repo pesrsonnel_repo.PersonnelRepository) PersonnelService {
	return PersonnelService{
		logger: logger,
		repo:   repo,
	}
}

func (s PersonnelService) CreateRadioTest(ctx context.Context, testModel personnel.RadioTest) (int, error) {
	var data = personnel.MapRequestRadioTestModelToServiceModel(testModel)
	id, err := s.repo.InsertRadioTest(ctx, data)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s PersonnelService) CreateTestText(ctx context.Context, testModel personnel.TextTest) (int, error) {
	var data = personnel.MapRequestTextTestModelToServiceModel(testModel)

	err := os.MkdirAll("./uploads/tests", os.ModePerm)
	if err != nil {
		return 0, err
	}

	actualFilePath := fmt.Sprintf("./uploads/tests/%d%s", time.Now().UnixNano(), filepath.Ext(testModel.Header.Filename))
	dst, err := os.Create(actualFilePath)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	_, err = io.Copy(dst, testModel.File)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.InsertTextTest(ctx, data, actualFilePath)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s PersonnelService) CreateRequest(ctx context.Context, reqModel personnel.Request) error {
	var data = personnel.MapRequestRequestModelToServiceRequestModel(reqModel)
	err := s.repo.InsertRequest(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (s PersonnelService) GetAllRequestsByUserId(ctx context.Context, userId string) ([]pesrsonnel_repo.Request, error) {
	getRequests, err := s.repo.GetAllRequests(ctx, userId)
	if err != nil {
		return []pesrsonnel_repo.Request{}, err
	}

	return getRequests, nil
}

func (s PersonnelService) GetTestByTestId(ctx context.Context, testId string) (personnel.RadioTest, error) {
	models, err := s.repo.GetTestById(ctx, testId)
	if err != nil {
		return personnel.RadioTest{}, err
	}
	var data = personnel.MapServiceRadioTestModelToRequestModel(models)

	return data, nil
}

func (s PersonnelService) AddOrganizations(ctx context.Context, request organization.AddOrganizationRequest) (int, error) {
	data := organization.MapAddOrganizationReqToModel(request)
	id, err := s.repo.InsertOrganization(ctx, data)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s PersonnelService) GetOrganizations(ctx context.Context) (model.GetOrganizationsModel, error) {
	organizations, err := s.repo.SelectOrganizations(ctx)
	if err != nil {
		return model.GetOrganizationsModel{}, err
	}

	return organizations, nil
}
