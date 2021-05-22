package transport

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/inconshreveable/log15"
	"github.com/stolarskis/goPlant/pkg/platform/pgsql"

	"github.com/stolarskis/goPlant/pkg/SensorData"
)

type Handler struct{}

func (h Handler) MoistureData(w http.ResponseWriter, r *http.Request) {
	err := uploadData(r, SensorData.MoistureSensor)
	if err != nil {
		log.Error(err.Error())
	}
}

func (h Handler) TempData(w http.ResponseWriter, r *http.Request) {
	err := uploadData(r, SensorData.TemperatureSensor)
	if err != nil {
		log.Error(err.Error())
	}
}

func (h Handler) lightData(w http.ResponseWriter, r *http.Request) {
	err := uploadData(r, SensorData.LightSensor)
	if err != nil {
		log.Error(err.Error())
	}
}

func uploadData(r *http.Request, sT SensorData.SensorType) error {
	sd, err := convertRequest(r)
	if err != nil {
		return err
	}

	sd.SensorType = sT

	pgsql.UploadData(sd)

	return nil
}

func convertRequest(r *http.Request) (SensorData.SensorData, error) {
	d, err := readBody(r.Body)
	if err != nil {
		return SensorData.SensorData{}, err
	}

	req, err := unmarshalData(d)
	if err != nil {
		return SensorData.SensorData{}, err
	}

	return req.convertToSensorData(), nil
}

func readBody(b io.ReadCloser) ([]byte, error) {
	res, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, errors.New("Moisture: Failed to read request body. " + err.Error())
	}

	return res, nil
}

func unmarshalData(b []byte) (DataUploadReq, error) {
	dReq := DataUploadReq{}
	err := json.Unmarshal(b, &dReq)
	if err != nil {
		return DataUploadReq{}, errors.New("Moisture: Failed to unmarshal request. " + err.Error())
	}

	return dReq, nil
}
