package config

import (
	"encoding/json"
	"io"
	"os"
)

type IConfig interface {
	Get() *Config
}

type Config struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinimumBalanceAmount int `json:"minimumBalanceAmount"`
}

var c = &Config{}

func (*Config) Get() *Config {
	file, err := os.Open(".config/" + env + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(read, c)
	if err != nil {
		panic(err)
	}

	return c
}

func NewConfig() IConfig {
	return c
}
