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
	err := uploadData(r, SensorData.SoilMoistureSensor)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Moisture data uploaded successfully"))
	}
}

func (h Handler) SoilTempData(w http.ResponseWriter, r *http.Request) {
	err := uploadData(r, SensorData.SoilTempSensor)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Soil temperature data uploaded successfully"))
	}
}

func (h Handler) AirTempData(w http.ResponseWriter, r *http.Request) {
	err := uploadData(r, SensorData.AirTempSensor)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		w.Write([]byte("Air temperature data uploaded successfully"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Air temperature data uploaded successfully"))
	}
}

func (h Handler) LightData(w http.ResponseWriter, r *http.Request) {
	err := uploadData(r, SensorData.LightSensor)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		w.Write([]byte("Humidity data uploaded successfully"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Humidity data uploaded successfully"))
	}
}

func (h Handler) HumidityData(w http.ResponseWriter, r *http.Request) {
	err := uploadData(r, SensorData.HumiditySensor)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Humidity data uploaded successfully"))
	}
}

func uploadData(r *http.Request, sT SensorData.SensorType) error {
	sd, err := convertRequest(r)
	if err != nil {
		return err
	}

	sd.SensorType = sT

	err = pgsql.UploadData(sd)
	if err != nil {
		return err
	}

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
