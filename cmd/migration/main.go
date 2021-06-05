package main

import (
	_ "github.com/lib/pq"
)

func main() {
	mc := migrationConig{
		dbName:     "goPlant",
		configPath: "",
	}

	MigrateDb(mc)
}
