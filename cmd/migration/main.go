package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/stolarskis/goPlant/utl/db"

	_ "github.com/lib/pq"
)

var tNames = []string{"moisture", "temperature", "light"}

const sname = "goPlant"
const isExists = "SELECT EXISTS(SELECT %s_name FROM information_schema.%s WHERE %s_name = '%s');"

func main() {

	db, err := db.New()
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
	}
}

func createTables(db *sql.DB) {

	const cT = "CREATE TABLE goplant.%s (id serial NOT NULL, value int4 NOT NULL, createdate date NULL DEFAULT CURRENT_DATE, recordtime date NULL, CONSTRAINT %s_pk PRIMARY KEY (id));"

	for _, t := range tNames {
		q := fmt.Sprintf(isExists, "table", "tables", "table", t)
		if !checkIfExists(db, q) {
			fmt.Println("Creating table: ", t)
			s := fmt.Sprintf(cT, t, t)
			_, err := db.Exec(s)
			if err != nil {
				log.Fatal(err)
			}

		}
	}
}

func checkIfExists(db *sql.DB, q string) bool {

	var e bool
	row, err := db.Query(q)
	defer row.Close()

	if err != nil {
		log.Fatal(err)
	}

	row.Next()
	row.Scan(&e)
	return e
}
