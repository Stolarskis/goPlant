package db

import (
	"database/sql"
	"fmt"
	"os"

	log "github.com/inconshreveable/log15"
)

type DbConfig struct {
	User string
	Pass string
	Name string
	Host string
}

var DbStg DbConfig

func New() (*sql.DB, error) {
	if (DbStg == DbConfig{}) {
		log.Crit("Database settings have not been initialized")
		os.Exit(1)
	}

	log.Debug(
		"Database User: " + DbStg.User +
			"\n Database Name: " + DbStg.Name +
			"\n Database Host: " + DbStg.Host +
			"\n Database Password: " + DbStg.Pass)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DbStg.User, DbStg.Pass, DbStg.Name, DbStg.Host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
