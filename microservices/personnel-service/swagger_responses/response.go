package swagger_responses

type CreateTestOkRes struct {
	Id      int    `json:"id" example:"1"`
	Message string `json:"error" example:""`
}

type CreateRequestOkRes struct {
	Message string `json:"error" example:""`
}

type GetRequestsOkRes struct {
	Error    string         `json:"error" example:""`
	Requests []RequestOkRes `json:"requests"`
}

type RequestOkRes struct {
	Name        string `json:"name" example:"Алексей"`
	Surname     string `json:"surname" example:"Тимошин"`
	TestId      int    `json:"test_id" example:"1"`
	Title       string `json:"title" example:"Вакансия frontend"`
	Description string `json:"description" example:"Next.js"`
}

type HTTPErrorCreateTest struct {
	Id      int    `json:"id" example:"0"`
	Message string `json:"error" example:"status bad request"`
}

type HTTPErrorCreateRequest struct {
	Message string `json:"error" example:"status bad request"`
}

type HTTPErrorGetRequests struct {
	Error    string `json:"error" example:"status bad request"`
	Requests []byte `json:"requests"`
}
