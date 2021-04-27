package db

import (
	"database/sql"
	"fmt"
	"log"
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
		log.Fatal("Database settings have not been initialized")
	}

	fmt.Println(DbStg.User, DbStg.Pass, DbStg.Name, DbStg.Host)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DbStg.User, DbStg.Pass, DbStg.Name, DbStg.Host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
