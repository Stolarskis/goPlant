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

type GetLatestSensorResp struct {
	Val float32
	Msg string
}

func (h Handler) PostMoistureData(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) PostSoilTempData(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) PostAirTempData(w http.ResponseWriter, r *http.Request) {
	err := uploadData(r, SensorData.AirTempSensor)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		w.Write([]byte("Error uploadin Air temperature data"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Air temperature data uploaded successfully"))
	}
}

func (h Handler) PostLightData(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) PostHumidityData(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) GetSensorReading(w http.ResponseWriter, r *http.Request) {
	sT, err := convertPathToSensorType(r.URL.Path)
	resp := GetLatestSensorResp{}
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusNotFound)
		resp.Msg = "Unable to determine Sensor Type"
	}

	v, err := getLatestSensorRead(sT)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		resp.Msg = "Unable to determine Sensor Type"
	}

	if v == -1 {
		log.Debug("Sensor " + sT.ToString() + "s table is empty")
		w.WriteHeader(http.StatusNotFound)
		resp.Msg = "There are no stored readings for " + sT.ToString() + " sensor type"
	} else {
		log.Debug("Retrieved latest reading of " + sT.ToString() + " sensor")
		w.WriteHeader(http.StatusOK)
		resp.Msg = "Successfully retrieved " + sT.ToString() + " sensor's latest reading"
		resp.Val = v
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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

func getLatestSensorRead(s SensorData.SensorType) (float32, error) {
	v, err := pgsql.GetLatestSensorReading(s)
	if err != nil {
		return 0, err
	}

	return v, nil
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
