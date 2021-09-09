package config

import (
	"errors"
	"os"
)

type DbConfig struct {
	User string
	Pass string
	Name string
	Host string
	Port string
}

type AppConfig struct {
	Host string
	Port string
}

func GetDbSettings() (DbConfig, error) {

	db := DbConfig{}

	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.Name = os.Getenv("DB_NAME")
	db.User = os.Getenv("DB_USERNAME")
	db.Pass = os.Getenv("DB_PASSWORD")

	err := checkDbSettings(db)
	if err != nil {
		return db, err
	}

	return db, nil
}

func GetAppSettings() (AppConfig, error) {
	app := AppConfig{}
	app.Host = os.Getenv("APP_HOST")
	app.Port = os.Getenv("APP_PORT")
	err := checkAppSettings(app)
	if err != nil {
		return app, err
	}

	return app, nil
}

func checkDbSettings(db DbConfig) error {
	if db.Host == "" || db.Name == "" || db.Pass == "" || db.Port == "" || db.User == "" {
		return errors.New("database setting(s) not found")
	}

	return nil
}

func checkAppSettings(app AppConfig) error {
	if app.Host == "" || app.Port == "" {
		return errors.New("application setting(s) not found")
	}

	return nil
}
