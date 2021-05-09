package pgsql

import (
	"database/sql"

	log "github.com/inconshreveable/log15"

	"github.com/stolarskis/goPlant/pkg/SensorData"
)

var DB *sql.DB

const ps = `INSERT INTO goplant.moisture
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

func UploadData(sd SensorData.SensorData) {

	log.Debug("Storing Data in Db." +
		"\n Sensor Type: " + sd.SensorType.ToString() +
		"\n Sensor Value: " + sd.Value)

	_, err := DB.Exec(ps, sd.Value, sd.RDate)
	if err != nil {
		log.Error(err.Error())
	}
}
