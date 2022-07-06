package serial

import (
	"github.com/thep0y/go-logger/log"
	"go.bug.st/serial"
)

type UsbSerial struct {
	Port      string
	BaudRate  int
	UsbObject serial.Port
}

func NewUsbSerial(Port string, BaudRate int) UsbSerial {
	usb := UsbSerial{Port: Port, BaudRate: BaudRate}
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		log.Infof("Found port: %v\n", port)
	}

	mode := &serial.Mode{
		BaudRate: usb.BaudRate,
	}
	port, err := serial.Open(usb.Port, mode)
	if err != nil {
		log.Fatal(err)
	}
	usb.UsbObject = port
	log.Infof("NewUsbSerial %s Initialized", usb.Port)
	return usb
}

func (usb *UsbSerial) Read(usbChan chan []byte) {
	log.Infof("USB Read Start %s", usb.Port)
	//buff := make([]byte, 100)
	buff := make([]byte, 100)
	for {
		n, err := usb.UsbObject.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			log.Infof("\nEOF")
			break
		}
		usbChan <- buff[:n]
		//log.Infof("%v", string(buff[:n]))
	}
}
