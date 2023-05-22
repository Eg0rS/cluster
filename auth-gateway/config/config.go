package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Read() *Settings {
	env, err := readEnv()
	fmt.Printf("Environment: %s\n", env)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var result Settings
	filePath := "./.config/" + string(env) + ".json"
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := json.Unmarshal(fileBytes, &result); err != nil {
		log.Fatalf("Unmarshalling error: %s", err)
		os.Exit(1)
	}

	return &result
}

func readEnv() (Env, error) {
	serviceEnv := os.Getenv(ServiceEnvVarName)
	if serviceEnv == "" {
		return EnvLocal, nil
	}
	switch Env(serviceEnv) {
	case Env(""):
		return EnvLocal, nil
	case EnvLocal:
		return EnvLocal, nil
	case EnvProd:
		return EnvProd, nil
	default:
		return "", fmt.Errorf("Unknown environment '%s'", os.Args[1])
	}
}
