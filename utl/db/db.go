package db

import (
	"database/sql"
	"fmt"
	"os"

	log "github.com/inconshreveable/log15"
	"github.com/stolarskis/goPlant/utl/config"
)

func New(configPath string) (*sql.DB, error) {

	dbC, _, err := config.GetDbSettings(configPath)
	if err != nil {
		log.Crit(err.Error())
	}

	if (dbC == config.DbConfig{}) {
		log.Crit("Database settings have not been initialized")
		os.Exit(1)
	}

	log.Debug(
		"Database User: " + dbC.User +
			"\n Database Name: " + dbC.Name +
			"\n Database Host: " + dbC.Host +
			"\n Database Password: " + dbC.Pass)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbC.User, dbC.Pass, dbC.Name, dbC.Host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
