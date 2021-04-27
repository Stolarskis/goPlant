package transport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
		fmt.Println("Moisture: Failed to read request body")
	}

	dReq := DataUploadReq{}
	err = json.Unmarshal(res, &dReq)
	if err != nil {
		fmt.Println("Moisture: Failed to unmarshal request")
	}

	sd := SensorData.SensorData(dReq)

	pgsql.UploadData(sd)
}

// func (h Handler) tempData(w http.ResponseWriter, r *http.Request) {}

// func (h Handler) lightData(w http.ResponseWriter, r *http.Request) {}
