package config

const ServiceEnvVarName = "SW_ENV"

const (
	EnvProd  = Env("prod")
	EnvLocal = Env("local")
)

type Env string
