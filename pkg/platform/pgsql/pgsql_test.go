package pgsql

import (
	"os"
	"strconv"
	"testing"

	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/stolarskis/goPlant/pkg/SensorData"
	"github.com/stolarskis/goPlant/utl/db"
)

var tSensorData = []SensorData.SensorData{
	{
		SensorId:   "1",
		PlantId:    "1",
		Value:      "10.123456789",
		RDate:      "1/23/21",
		SensorType: SensorData.LightSensor,
	},
	{
		SensorId:   "1",
		PlantId:    "1",
		Value:      "64.93775912",
		RDate:      "1/23/21",
		SensorType: SensorData.AirTempSensor,
	},
	{
		SensorId:   "1",
		PlantId:    "1",
		Value:      "92.09574212",
		RDate:      "1/23/21",
		SensorType: SensorData.SoilMoistureSensor,
	},
	{
		SensorId:   "1",
		PlantId:    "1",
		Value:      "95.16302673",
		RDate:      "1/23/21",
		SensorType: SensorData.SoilTempSensor,
	},
	{
		SensorId:   "1",
		PlantId:    "1",
		Value:      "56.69361860",
		RDate:      "1/23/21",
		SensorType: SensorData.HumiditySensor,
	},
}

func TestMain(m *testing.M) {
	err := crDatabaseConn()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	if DB == nil {
		log.Error("Database connection is nil")
		os.Exit(1)
	}
	err = DB.Ping()
	if err != nil {
		log.Error("Ping to database failed" + "\n" + err.Error())
		os.Exit(1)
	}

	m.Run()
}

func TestUpload(t *testing.T) {

	for _, s := range tSensorData {
		err := UploadData(s)
		if err != nil {
			t.Errorf("Test Upload: Failed to upload data to database." + "\n" + err.Error())
		}

		r, err := GetLatestSensorReading(s.SensorType)
		if err != nil {
			t.Errorf("Test Upload: Failed to get data from latest upload" + "\n" + err.Error())
		}

		sVal, err := convStrFloat(s.Value)
		if err != nil {
			t.Errorf("Failed to convert string to float32" + "\n" + err.Error())
		}

		if r != sVal {
			t.Errorf("Test Upload: Value returned from database doesn't match value uploaded.")
		}
	}
}

func crDatabaseConn() error {
	db, err := db.New("../../../env/local/")
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func convStrFloat(v string) (float32, error) {

	r, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0, err
	}

	return float32(r), nil
}
