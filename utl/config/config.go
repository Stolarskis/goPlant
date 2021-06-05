package config

import (
	"errors"

	log "github.com/inconshreveable/log15"
	"github.com/spf13/viper"
)

type DbConfig struct {
	User string
	Pass string
	Name string
	Host string
}

type SvcConfig struct {
	Host string
	Port string
}

func GetDbSettings(configPath string) (DbConfig, SvcConfig, error) {

	path, err := getConfigPath(configPath)
	if err != nil {
		log.Warn(err.Error())
	}

	setViperStg(path)

	dbC, sC, err := readConf()

	if err != nil {
		return DbConfig{}, SvcConfig{}, err
	}

	return dbC, sC, nil

}

func setViperStg(env string) {
	viper.AddConfigPath(env)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
}

func readConf() (DbConfig, SvcConfig, error) {
	dbC := DbConfig{}
	sC := SvcConfig{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return dbC, sC, errors.New("Failed to read config file")
	}

	viper.UnmarshalKey("db", &dbC)
	viper.UnmarshalKey("server", &sC)

	return dbC, sC, nil
}

func getConfigPath(cp string) (string, error) {

	if cp == "" {
		return "env/local/", errors.New("ConfigPath is empty, using default local config")
	}

	return cp, nil
}
