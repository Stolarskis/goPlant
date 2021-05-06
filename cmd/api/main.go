package main

import (
	"os"

	log "github.com/inconshreveable/log15"
	"github.com/stolarskis/goPlant/pkg/platform/pgsql"
	"github.com/stolarskis/goPlant/utl/db"
	"github.com/stolarskis/goPlant/utl/server"

	_ "github.com/lib/pq"
	"goji.io"
)

func main() {

	db, err := db.New()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	pgsql.DB = db
	m := goji.NewMux()

	server.NewHTTP(m)
	err = server.Start(m)
	if err != nil {
		log.Error(err.Error())
	}
}
