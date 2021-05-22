package SensorData

type SensorType int

const (
	MoistureSensor    SensorType = iota
	TemperatureSensor SensorType = iota
	LightSensor       SensorType = iota
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
	case MoistureSensor:
		return "Moisture"
	case TemperatureSensor:
		return "Temperature"
	case LightSensor:
		return "Light"
	}

	return ""
}
