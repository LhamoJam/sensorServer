package message

type SensorMessage struct {
	Message
	Sensortype string `json:"sensortype"`
}

func (msg SensorMessage) GetContent() []byte {

	return msg.Content
}
