package personnel

import "personnel_service/database/pesrsonnel_repo"

type CreateTestResponse struct {
	Id    int    `json:"id"`
	Error string `json:"error"`
}

type CreateRequestResponse struct {
	Error string `json:"error"`
}

type GetRequestsResponse struct {
	Requests []pesrsonnel_repo.Request `json:"requests"`
	Error    string                    `json:"error"`
}
