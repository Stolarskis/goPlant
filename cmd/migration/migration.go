package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/stolarskis/goPlant/pkg/platform/pgsql"

	log "github.com/inconshreveable/log15"

	"github.com/stolarskis/goPlant/utl/db"

	_ "github.com/lib/pq"
)

type migrationConig struct {
	dbName     string
	configPath string
}

const sname = "goPlant"
const isExists = "SELECT EXISTS(SELECT %s_name FROM information_schema.%s WHERE %s_name = '%s');"

var TableNames = pgsql.TableNames

func MigrateDb(config migrationConig) {

	db, err := db.New(config.configPath)
	if err != nil {
		log.Crit("Failed to create db connection: " + err.Error())
		os.Exit(1)
	}

	createSchema(db)
	createTables(db)
}

func createSchema(db *sql.DB) {
	//Query to check if schema exists
	c := fmt.Sprintf(isExists, "schema", "schemata", "schema", "goplant")

	//If the schema doesn't exist, create new schema
	if !checkIfExists(db, c) {
		_, err := db.Exec("CREATE SCHEMA \"goplant\"")
		if err != nil {
			log.Error("Failed to create goplant schema: " + err.Error())
		}
	}
}

func createTables(db *sql.DB) {

	const cT = "CREATE TABLE goplant.%s (id serial NOT NULL, value int4 NOT NULL, recordtime timestamp NULL, createdate timestamp NULL DEFAULT CURRENT_DATE, CONSTRAINT %s_pk PRIMARY KEY (id));"

	for _, t := range TableNames {
		q := fmt.Sprintf(isExists, "table", "tables", "table", t)
		if !checkIfExists(db, q) {
			log.Debug("Creating Table " + t)
			s := fmt.Sprintf(cT, t, t)
			_, err := db.Exec(s)
			if err != nil {
				log.Error("Failed to create sensor data tables: " + err.Error())
				os.Exit(1)
			}

		}
	}
}

func checkIfExists(db *sql.DB, q string) bool {

	var e bool
	row, err := db.Query(q)
	defer row.Close()

	if err != nil {
		log.Error("Failed to check if schema already exists: " + err.Error())
		os.Exit(1)
	}

	row.Next()
	row.Scan(&e)
	return e
}
