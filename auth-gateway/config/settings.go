package config

type Settings struct {
	Port                    int    `json:"port"`
	TokenSecret             string `json:"token_secret"`
	ServerTimeout           int64  `json:"server_timeout_sec"`
	RequestForwarderTimeout int64  `json:"request_forwarder_timeout_sec"`
	JWTHeader               string `json:"jwt_header"`
	AuthService             string `json:"auth_service"`
}
