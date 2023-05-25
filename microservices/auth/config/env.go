package config

const ServiceEnvVarName = "env"

const (
	EnvProd  = Env("prod")
	EnvLocal = Env("local")
)

type Env string
