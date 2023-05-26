package organization

type AddOrganizationResponse struct {
	Id    int    `json:"id" example:"1"`
	Error string `json:"error" example:""`
}

type AddOrganizationBadResponse struct {
	Id    int    `json:"id" example:"0"`
	Error string `json:"error" example:"bad request"`
}

type OrgInfoResponse struct {
	Name    string `json:"name" example:"Amazon"`
	Address string `json:"address" example:"st. Washington Jonson street h.27"`
}

type OrganizationsInfoResponse struct {
	Organizations []OrgInfoResponse `json:"organizations"`
}
