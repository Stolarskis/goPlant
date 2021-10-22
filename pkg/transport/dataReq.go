package transport

import (
	"errors"
	"strings"

	"github.com/stolarskis/goPlant/pkg/SensorData"
)

type DataUploadReq struct {
	SensorId string `json:"id"`
	PlantId  string `json:"plantId"`
	Value    string `json:"value"`
}

var endpointToSensorType = map[string]SensorData.SensorType{
	"moisture": SensorData.SoilMoistureSensor,
	"soilTemp": SensorData.SoilTempSensor,
	"airTemp":  SensorData.AirTempSensor,
	"light":    SensorData.LightSensor,
	"humidity": SensorData.HumiditySensor,
}

func (d DataUploadReq) convertToSensorData() SensorData.SensorData {

	sd := SensorData.SensorData{}

	sd.SensorId = d.SensorId
	sd.PlantId = d.PlantId
	sd.Value = d.Value

	return sd
}

func convertPathToSensorType(p string) (SensorData.SensorType, error) {

	s := strings.Split(p, "/")

	v, ok := endpointToSensorType[s[2]]
	if ok && v != SensorData.UnknownSensor {
		return v, nil
	} else {
		return SensorData.UnknownSensor, errors.New("failed to determine sensor Type")
	}
}
