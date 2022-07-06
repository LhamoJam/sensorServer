package sensor

import (
	"github.com/thep0y/go-logger/log"
	"time"
)

type TempSensor struct {
	Sensor
}

func NewTempSensor(pin int) *TempSensor {
	sensor := &TempSensor{Sensor{Name: "TEMP", Pin: pin}}
	return sensor
}

func (sensor TempSensor) SensorInit() {
	log.Infof("%s SensorInit", sensor.Name)
}

func (sensor TempSensor) GetData(msgChan chan []byte) {
	for {
		sensor.Sensor.Message.Content = []byte("Fuck your 2!")
		time.Sleep(1 * time.Second)
		msgChan <- sensor.Sensor.Message.Content
	}
}
func (sensor TempSensor) Opration(recvChan chan []byte) {
	for {
		for data := range recvChan {
			switch string(data) {
			case "123":
				log.Infof("%s Sensor COMMAND: %v", sensor.Name, data)
			}
		}
	}
}
