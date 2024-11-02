package config

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Development struct {
		Dialect  string
		Database string
		Username string
		Password string
		Host     string
		Port     int
	}
}

var DbConfig DatabaseConfig

func LoadConfig() {
	data, err := os.ReadFile("config/database.yml")
	if err != nil {
		panic("Failed to read config file:" + err.Error())
	}

	err = yaml.Unmarshal(data, &DbConfig)
	if err != nil {
		panic("Failed to parse config file:" + err.Error())
	}
}

func GetDatabaseConfig() string {
	return DbConfig.Development.Username + ":" + DbConfig.Development.Password + "@" +
		"tcp(" + DbConfig.Development.Host + ":" + strconv.Itoa(DbConfig.Development.Port) + ")" +
		"/" + DbConfig.Development.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}
