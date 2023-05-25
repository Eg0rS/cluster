package config

type Settings struct {
	Port               int    `json:"port"`
	DbConnectionString string `json:"db_connection_string"`
	ClientSecret       string `json:"client_secret"`
	JwtSecret          string `json:"jwt_secret"`
	RefreshTokenSecret string `json:"refresh_token_secret"`
	AccessTokenTTL     int64  `json:"access_token_ttl"`
	RefreshTokenTTL    int64  `json:"refresh_token_ttl"`
	SuperPassword      string `json:"super_password"`
}
