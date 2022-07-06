package server

import (
	"bytes"
	"github.com/gorilla/websocket"
	"github.com/thep0y/go-logger/log"
	"net/http"
	"sensorServer/sensor"
)

func TdsSensorHandler(w http.ResponseWriter, r *http.Request) {
	//msgChan := &usbChan
	recvChan := make(chan []byte, 5)
	// 初始化传感器

	sensorObj := sensor.NewTdsSensor(12)
	sensorObj.SensorInit()
	//go sensor.GetData(msgChan)
	go sensorObj.Opration(recvChan)

	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warnf("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	go func() {
		for {
			// 接收数据
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Warnf("Error during message reading:", err)
				break
			}
			log.Infof("Received: %s", msg)
			recvChan <- msg
		}
	}()
	for {
		data := <-usbChan
		res := tdsdataprocessing(&data)
		// 发送数据
		err = conn.WriteMessage(websocket.BinaryMessage, res)
		if err != nil {
			log.Warnf("Error during message writing:", err)
			break
		}
	}
}

func tdsdataprocessing(data *[]byte) []byte {
	temp := bytes.Split(*data, []byte(","))
	return temp[0]
}
