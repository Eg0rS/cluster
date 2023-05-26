package profile

import "profile_service/model"

func MapClientToServiceUpsertInfo(req UpsertUserInfoReq) model.UpsertUserInfoModel {
	return model.UpsertUserInfoModel{
		FirstName:  req.FirstName,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
		City:       req.City,
		University: req.University,
		Age:        req.Age,
		Education:  req.Education,
	}
}
