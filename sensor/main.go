package sensor

import (
	"sensorServer/message"
)

type Sensor struct {
	Name    string
	Pin     int
	Message message.SensorMessage
}

type SensorOpration interface {
	SensorInit() error
	GetData(msgchan chan []byte)
	Opration(recvchan chan []byte)
}
