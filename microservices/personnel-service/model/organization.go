package model

type AddOrganizationModel struct {
	Name             string
	Address          string
	FirstCoordinate  float64
	SecondCoordinate float64
}

type OrganizationInfo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GetOrganizationsModel struct {
	Organizations []OrganizationInfo `json:"organizations"`
}
