package server

import (
	"github.com/gorilla/websocket"
	"github.com/thep0y/go-logger/log"
	"net/http"
	"sensorServer/serial"
)

func UsbSerialHandler(w http.ResponseWriter, r *http.Request) {
	// USB 串口
	var usb = serial.NewUsbSerial("/dev/ttyUSB0", 9600)
	usbChan := make(chan []byte, 100)
	//recvChan := make(chan []byte, 5)
	// 初始化传感器

	go usb.Read(usbChan)
	//go sensor.Opration(recvChan)

	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warnf("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	//go func() {
	//	for {
	//		// 接收数据
	//		_, msg, err := conn.ReadMessage()
	//		if err != nil {
	//			log.Warnf("Error during message reading:", err)
	//			break
	//		}
	//		log.Infof("Received: %s", msg)
	//		recvChan <- msg
	//	}
	//}()
	for {
		// 发送数据
		err = conn.WriteMessage(websocket.BinaryMessage, <-usbChan)
		if err != nil {
			log.Warnf("Error during message writing:", err)
			break
		}
	}
}
