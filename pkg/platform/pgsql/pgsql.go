package pgsql

import (
	"database/sql"
	"errors"

	log "github.com/inconshreveable/log15"

	"github.com/stolarskis/goPlant/pkg/SensorData"
)

var DB *sql.DB

var TableNames = map[SensorData.SensorType]string{
	SensorData.SoilMoistureSensor: "soilmoisture",
	SensorData.SoilTempSensor:     "soiltemperature",
	SensorData.AirTempSensor:      "airtemperature",
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

var getQueries = map[SensorData.SensorType]string{
	SensorData.SoilMoistureSensor: mGetLastVal,
	SensorData.SoilTempSensor:     stGetLastVal,
	SensorData.AirTempSensor:      atGetLastVal,
	SensorData.LightSensor:        lGetLastVal,
	SensorData.HumiditySensor:     hGetLastVal,
}

func UploadData(sd SensorData.SensorData) error {

	log.Debug("Storing Data into Table: " + TableNames[sd.SensorType] +
		"\n Sensor Type: " + sd.SensorType.ToString() +
		"\n Sensor Value: " + sd.Value)

	_, err := DB.Exec(uploadQueries[sd.SensorType], sd.Value, sd.RDate)
	if err != nil {
		return err
	}
	return nil
}

func GetLatestSensorReading(s SensorData.SensorType) (float32, error) {
	log.Debug("Grabbing latest reading of " + s.ToString())

	r, err := DB.Query(getQueries[s])
	if err != nil {
		return 0, err
	}

	v, err := getFirstRowVal(r)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func getFirstRowVal(r *sql.Rows) (float32, error) {
	var v *float32 = new(float32)
	e := r.Next()
	if !e {
		return 0, errors.New("unable to get sensor value, result set is empty")
	}
	err := r.Scan(v)
	if err != nil {
		return 0, err
	}

	return *v, nil
}
