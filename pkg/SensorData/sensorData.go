package SensorData

//Defines 3 sensor types, moisture, temperature, and light
const (
	mSensor = iota
	tSensor = iota
	lSensor = iota
)

type SensorData struct {
	SensorId   string
	PlantId    string
	SensorType string
	Value      string
	RDate      string
}
