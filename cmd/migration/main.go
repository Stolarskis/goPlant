package main

import (
	_ "github.com/lib/pq"
)

func main() {
	mc := migrationConfig{
		dbName:     "goPlant",
		configPath: "",
	}

	MigrateDb(mc)
}
