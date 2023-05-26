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
	X       float64
	Y       float64
}

type GetOrganizationsModel struct {
	Organizations []OrganizationInfo `json:"organizations"`
}
