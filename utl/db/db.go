package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func New() (*sql.DB, error) {

	dbU := os.Getenv("DB_USER")
	dbP := os.Getenv("DB_PASSWORD")
	dbN := os.Getenv("DB_NAME")
	dbH := os.Getenv("DB_HOST")

	if dbU == "" || dbP == "" || dbN == "" {
		return nil, errors.New("Database env var not set")
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbU, dbP, dbN, dbH)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil

}
