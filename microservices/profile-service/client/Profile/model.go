package profile

type UpsertUserInfoReq struct {
	Surname    string `json:"surname" example:"Путин"`
	FirstName  string `json:"firstname" example:"Владимир"`
	Patronymic string `json:"patronymic" example:"Владимирович"`
	City       string `json:"city" example:"Москва"`
	University string `json:"university" example:"ИТМО"`
	Age        int    `json:"age" example:"68"`
	Education  string `json:"education" example:"Бакалавриат"`
}
