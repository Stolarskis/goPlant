package transport

import "github.com/stolarskis/goPlant/pkg/SensorData"

type DataUploadReq struct {
	SensorId string `json:"id"`
	PlantId  string `json:"plantId"`
	Value    string `json:"value"`
	RDate    string `json:"recordDate"`
}

func (d DataUploadReq) convertToSensorData() SensorData.SensorData {

	sd := SensorData.SensorData{}

	sd.SensorId = d.SensorId
	sd.PlantId = d.PlantId
	sd.Value = d.Value
	sd.RDate = d.RDate

	return sd

}
