package organization

type AddOrganizationRequest struct {
	Name             string  `json:"name"`
	Address          string  `json:"address"`
	FirstCoordinate  float64 `json:"first_coordinate"`
	SecondCoordinate float64 `json:"second_coordinate"`
}
