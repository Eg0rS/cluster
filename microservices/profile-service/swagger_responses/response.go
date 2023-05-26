package profile_responses

type UpsertGoodResponse struct {
	Error string `json:"error" example:""`
}

type UpsertBadResponse struct {
	Error string `json:"error" example:"Bad request, error on service"`
}
