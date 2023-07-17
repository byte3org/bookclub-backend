package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config : struct
type Config struct {
	Environment         string `json:"env"`
	Port                int    `json:"port"`
	Db                  string `json:"db"`
	TokenSecretKey      string `json:"token_secret_key"`
	AdminTokenSecretKey string `json:"admin_token_secret_key"`
}

var config *Config

func initialize() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("[x] error : ", err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("[x] error", err.Error())
	}
	config = &conf
}

func GetConfig() *Config {
	return config
}
