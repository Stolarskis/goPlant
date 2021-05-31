package main

import (
	"os"

	"github.com/stolarskis/goPlant/utl/config"

	log "github.com/inconshreveable/log15"
	"github.com/stolarskis/goPlant/pkg/platform/pgsql"
	"github.com/stolarskis/goPlant/utl/db"
	"github.com/stolarskis/goPlant/utl/server"

	_ "github.com/lib/pq"
	"goji.io"
)

func main() {

	dbC, sC, err := config.GetDbSettings()
	if err != nil {
		log.Error(err.Error())
	}

	db.DbStg = dbC
	server.ServerStg = sC

	db, err := db.New()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	pgsql.DB = db

	m := goji.NewMux()
	server.NewHTTP(m)
	server.Start(m)
	if err != nil {
		log.Error(err.Error())
	}
}
