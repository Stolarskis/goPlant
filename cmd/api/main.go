package main

import (
	"log"

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
		log.Fatal("Failed to read config file")
	}
	viper.UnmarshalKey("db", &dbC)
	viper.UnmarshalKey("server", &sC)

	db.DbStg = dbC
	server.ServerStg = sC

	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	pgsql.DB = db

	m := goji.NewMux()
	server.NewHTTP(m)
	server.Start(m)
}
