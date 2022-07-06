package server

import (
	"github.com/gorilla/websocket"
	"github.com/thep0y/go-logger/log"
	"net/http"
	"sensorServer/serial"
)

// tcp升级websocket
var upgrader = websocket.Upgrader{} // use default options

// USB 串口
var usb = serial.NewUsbSerial("/dev/ttyUSB0", 9600)
var usbChan = make(chan []byte, 100)

func ServerInitial(port string) {
	go usb.Read(usbChan)
	http.HandleFunc("/tds", TdsSensorHandler)
	http.HandleFunc("/temp", TempSensorHandler)
	//http.HandleFunc("/UsbSerial", UsbSerialHandler)
	log.Infof("ServerStart!!!")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
