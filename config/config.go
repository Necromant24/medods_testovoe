package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type AppConfig struct {
	DbConnectionString string `json: "dbConnectionString"`
	HostPort           int    `json: "hostPort"`
}

var appConfig *AppConfig

func GetConfig() *AppConfig {
	return appConfig
}

func LoadConfiguration() {
	var settings *AppConfig = &AppConfig{}
	configFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&settings); err != nil {
		fmt.Println("parsing config file", err.Error())
	}

	appConfig = settings

	fmt.Println("db conn str - "+settings.DbConnectionString, "host port - "+strconv.Itoa(settings.HostPort))

}
