package pgsql

import (
	"os"
	"testing"

	log "github.com/inconshreveable/log15"
	_ "github.com/lib/pq"
	"github.com/stolarskis/goPlant/pkg/SensorData"
	"github.com/stolarskis/goPlant/utl/db"
)

var tSensorData = []SensorData.SensorData{
	{
		SensorId:   "1",
		PlantId:    "1",
		Value:      "10",
		RDate:      "1/23/21",
		SensorType: SensorData.LightSensor,
	},
}

func TestUpload(t *testing.T) {

	db, err := db.New("../../../env/test/")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	DB = db

	for _, v := range tSensorData {
		err := UploadData(v)
		if err != nil {
			t.Log("Test Upload: Failed to upload data to database.")
			t.Fail()
		}
	}
}
