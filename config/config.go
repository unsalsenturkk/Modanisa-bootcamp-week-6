package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type IConfig interface {
	Get() *Config
}

type Config struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinimumBalanceAmount int `json:"minimumBalanceAmount"`
}

var c = &Config{}

func init() {

	mydir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf(mydir)

	if strings.Contains(mydir, "service") {
		os.Chdir("..")
	} else if strings.Contains(mydir, "controller") {
		os.Chdir("..")
	}

	file, err := os.Open(".config\\" + env + ".json")
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
}

func (*Config) Get() *Config {
	return c
}

func NewConfig() IConfig {
	return c
}
