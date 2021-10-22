package pgsql

import (
	"database/sql"
	"errors"
	"math"
	"strconv"
	"time"

	log "github.com/inconshreveable/log15"

	"github.com/stolarskis/goPlant/pkg/SensorData"
)

const tFormat = "02 Jan 06 15:04:01 MST"

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

	v, err := roundSensorVal(sd.Value, 0.001, 3)
	if err != nil {
		return err
	}

	log.Debug("Storing Data into Table: " + TableNames[sd.SensorType] +
		"\n Sensor Type: " + sd.SensorType.ToString() +
		"\n Sensor Value: " + v)

	_, err = DB.Exec(uploadQueries[sd.SensorType], v, time.Now().Format(tFormat))
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

func roundSensorVal(v string, unit float64, prec int) (string, error) {
	r, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return "", err
	}
	r = math.Round(r/unit) * unit

	rS := strconv.FormatFloat(r, 'f', prec, 64)
	if rS == "" {
		return "", errors.New("pgSql: Failed to convert rounded sensor value back into string")
	}

	return rS, nil
}
