package SensorData

type SensorType int

const (
	SoilMoistureSensor SensorType = iota
	SoilTempSensor     SensorType = iota
	AirTempSensor      SensorType = iota
	HumiditySensor     SensorType = iota
	LightSensor        SensorType = iota
)

type SensorData struct {
	SensorId   string
	PlantId    string
	Value      string
	RDate      string
	SensorType SensorType
}

func (sT SensorType) ToString() string {
	switch sT {
	case SoilMoistureSensor:
		return "Soil Moisture"
	case SoilTempSensor:
		return "Soil Temperature"
	case AirTempSensor:
		return "Air Temperature"
	case HumiditySensor:
		return "Air Humidity"
	case LightSensor:
		return "Light"
	}

	return ""
}
