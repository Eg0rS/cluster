package organization

import "personnel_service/model"

func MapAddOrganizationReqToModel(requestModel AddOrganizationRequest) model.AddOrganizationModel {
	return model.AddOrganizationModel{
		Name:             requestModel.Name,
		Address:          requestModel.Address,
		FirstCoordinate:  requestModel.FirstCoordinate,
		SecondCoordinate: requestModel.SecondCoordinate,
	}
}
