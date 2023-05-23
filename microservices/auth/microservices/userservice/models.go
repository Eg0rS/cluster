package userservice

type UserInfo struct {
	UserName     string  `json:"UserName"`
	EmployeeLink int     `json:"EmployeeLink"`
	Rights       *Rights `json:"Rights,omitempty"`
}

type Rights struct {
	BuyTripAccount int `json:"BuyTripAccount"`
}
