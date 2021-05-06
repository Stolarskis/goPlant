package transport

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/inconshreveable/log15"

	"github.com/stolarskis/goPlant/pkg/platform/pgsql"

	"github.com/stolarskis/goPlant/pkg/SensorData"
)

type Handler struct{}

type DataUploadReq struct {
	SensorId   string `json:"id"`
	PlantId    string `json:"plantId"`
	SensorType string `json:"sensorType"`
	Value      string `json:"value"`
	RDate      string `json:"recordDate"`
}

func (h Handler) MoistureData(w http.ResponseWriter, r *http.Request) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Moisture: Failed to read request body. " + err.Error())
	}

	dReq := DataUploadReq{}
	err = json.Unmarshal(res, &dReq)
	if err != nil {
		log.Error("Moisture: Failed to unmarshal request. " + err.Error())
	}

	sd := SensorData.SensorData(dReq)

	pgsql.UploadData(sd)
}

// func (h Handler) tempData(w http.ResponseWriter, r *http.Request) {}

// func (h Handler) lightData(w http.ResponseWriter, r *http.Request) {}
