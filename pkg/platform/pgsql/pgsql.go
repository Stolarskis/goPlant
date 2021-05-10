package pgsql

import (
	"database/sql"

	log "github.com/inconshreveable/log15"

	"github.com/stolarskis/goPlant/pkg/SensorData"
)

var DB *sql.DB

var TableNames = map[SensorData.SensorType]string{
	SensorData.SoilMoistureSensor: "soilMoisture",
	SensorData.SoilTempSensor:     "soilTemperature",
	SensorData.AirTempSensor:      "airTemperature",
	SensorData.LightSensor:        "light",
	SensorData.HumiditySensor:     "humidity",
}

var uploadQueries = map[SensorData.SensorType]string{
	SensorData.SoilMoistureSensor: mAddData,
	SensorData.SoilTempSensor:     sTAddData,
	SensorData.AirTempSensor:      aTAddData,
	SensorData.LightSensor:        lAddData,
	SensorData.HumiditySensor:     hAddData,
}

func UploadData(sd SensorData.SensorData) {

	log.Debug("Storing Data into Table: " + TableNames[sd.SensorType] +
		"\n Sensor Type: " + sd.SensorType.ToString() +
		"\n Sensor Value: " + sd.Value)

	_, err := DB.Exec(uploadQueries[sd.SensorType], sd.Value, sd.RDate)
	if err != nil {
		log.Error(err.Error())
	}
}
