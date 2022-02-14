package config

import "os"

const (
	envKey   = "APP_ENV"
	envLocal = "local"
	envProd  = "prod"
)

var env = GetEnv(envKey, envLocal)

func Env() string {
	return env
}

func GetEnv(key, def string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	return def
}
