package main

import (
	"log"

	"github.com/stolarskis/goPlant/pkg/platform/pgsql"
	"github.com/stolarskis/goPlant/utl/db"
	"github.com/stolarskis/goPlant/utl/server"

	_ "github.com/lib/pq"
	"goji.io"
)

func main() {

	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	pgsql.DB = db
	m := goji.NewMux()

	server.NewHTTP(m)
	server.Start(m)
}
