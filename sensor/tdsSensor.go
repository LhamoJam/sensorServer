package sensor

import (
	"github.com/thep0y/go-logger/log"
	"time"
)

type TdsSensor struct {
	Sensor
}

func NewTdsSensor(pin int) *TdsSensor {
	sensor := &TdsSensor{Sensor{Name: "TDS", Pin: pin}}
	return sensor
}

func (sensor TdsSensor) SensorInit() {
	log.Infof("%s SensorInit", sensor.Name)
}

func (sensor TdsSensor) GetData(msgChan chan []byte) {
	for {
		sensor.Sensor.Message.Content = []byte("Fuck your pussy!")
		time.Sleep(1 * time.Second)
		msgChan <- sensor.Sensor.Message.Content
	}
}
func (sensor TdsSensor) Opration(recvChan chan []byte) {
	for {
		for data := range recvChan {
			switch string(data) {
			case "123":
				log.Infof("%s Sensor COMMAND: %v", sensor.Name, data)
			}
		}
	}
}
