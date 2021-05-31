package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
	"github.com/stolarskis/goPlant/utl/db"
	"github.com/stolarskis/goPlant/utl/server"
)

func GetDbSettings() (db.DbConfig, server.ServerConfig, error) {

	env, err := getEnvPath()
	if err != nil {
		return db.DbConfig{}, server.ServerConfig{}, err
	}

	setViperStg(env)

	dbC, sC, err := readConf()

	if err != nil {

		return db.DbConfig{}, server.ServerConfig{}, err
	}

	return dbC, sC, nil

}

func setViperStg(env string) {
	viper.AddConfigPath(env)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
}

func readConf() (db.DbConfig, server.ServerConfig, error) {
	dbC := db.DbConfig{}
	sC := server.ServerConfig{}

	err := viper.ReadInConfig()
	if err != nil {
		return dbC, sC, errors.New("Failed to read config file")
	}

	viper.UnmarshalKey("db", &dbC)
	viper.UnmarshalKey("server", &sC)

	return dbC, sC, nil
}

func getEnvPath() (string, error) {
	env := os.Getenv("CONFIG_ENV")

	if env == "" {
		return "", errors.New("Environment config environment variable, \"CONFIG_ENV\" is not set. \n" +
			"To set Environment config set environment variable CONFIG_ENV to a config environment. \n" +
			"The names of the environments are the names of the directories that contain them, such as \"local\" or \"test\"")
	}

	path := "./env/" + env + "/"

	return path, nil
}
