package Serial

import (
	"fmt"
	"log"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

func PrintPorts() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}

func PrintDetailedPorts() {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		return
	}
	for _, port := range ports {
		fmt.Printf("Found port: %s\n", port.Name)
		if port.IsUSB {
			fmt.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("   USB serial %s\n", port.SerialNumber)
		}
	}
}

func GetArduinos() []Arduino {
	ports, err := enumerator.GetDetailedPortsList()
	var arduino []Arduino
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No Serialports found!")
	}
	for _, port := range ports {
		if port.IsUSB {
			fmt.Printf("Serial USB port: %v\n", port)
			arduino = append(arduino, Arduino{port: port.Name})
		}
	}

	return arduino
}
