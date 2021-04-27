package pgsql

import (
	"database/sql"
	"fmt"

	"github.com/stolarskis/goPlant/pkg/SensorData"
)

var DB *sql.DB

func UploadData(sd SensorData.SensorData) {
	fmt.Println(sd)

	const ps = `INSERT INTO goplant.moisture
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

	_, err := DB.Exec(ps, sd.Value, sd.RDate)
	if err != nil {
		fmt.Println(err.Error())
	}
}
