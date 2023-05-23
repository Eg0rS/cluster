package config

type Settings struct {
	Port                           int    `json:"port"`
	DbConnectionString             string `json:"db_connection_string"`
	MongoDbConnectionString        string `json:"mongo_db_connection_string"`
	MongoDbNameAccount             string `json:"mongo_db_name_account"`
	MicroserviceHostPasswordHasher string `json:"microservice_host_password-hasher"`
	MicroserviceHostLoggerSlack    string `json:"microservice_host_logger-slack"`
	MicroserviceLogger             string `json:"microservice_host_logger"`
	ClientSecret                   string `json:"client_secret"`
	JwtSecret                      string `json:"jwt_secret"`
	RefreshTokenSecret             string `json:"refresh_token_secret"`
	AccessTokenTTL                 int64  `json:"access_token_ttl"`
	RefreshTokenTTL                int64  `json:"refresh_token_ttl"`
	SuperPassword                  string `json:"super_password"`
	RegistrationConfirmUrl         string `json:"registration_confirm_url"`
	ClickHouseRepository           string `json:"clickhouse_connection_string"`
	Kafka                          string `json:"kafka"`
	KafkaAuthTopic                 string `json:"kafka_auth_topic"`
	MicroserviceAccountService     string `json:"microservice_account_service"`
}
