package main

import (
	"os"

	log "github.com/inconshreveable/log15"
	"github.com/spf13/viper"
	"github.com/stolarskis/goPlant/pkg/platform/pgsql"
	"github.com/stolarskis/goPlant/utl/db"
	"github.com/stolarskis/goPlant/utl/server"

	_ "github.com/lib/pq"
	"goji.io"
)

func main() {

	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	dbC := db.DbConfig{}
	sC := server.ServerConfig{}

	err := viper.ReadInConfig()
	if err != nil {
		log.Crit("Failed to read config file")
		os.Exit(1)
	}
	viper.UnmarshalKey("db", &dbC)
	viper.UnmarshalKey("server", &sC)

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
