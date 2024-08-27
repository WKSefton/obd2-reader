package main

import (
	"fmt"
	"log"

	"github.com/WKSefton/obd2-reader/obd"
	"github.com/WKSefton/obd2-reader/serial"
)

func main() {
	// Open the serial port
	port, err := serial.OpenSerialPort("COM3") // Adjust for your COM port
	if err != nil {
		log.Fatalf("Failed to open serial port: %v", err)
	}
	defer serial.CloseSerialPort(port)

	// Initialize the ELM327 device
	obd.InitializeELM327(port)

	// Send a command to read the vehicle speed
	speed := obd.SendOBDCommand("010D", port)
	fmt.Printf("Vehicle Speed: %s km/h\n", speed)

	// Send a command to read the engine RPM
	rpm := obd.SendOBDCommand("010C", port)
	fmt.Printf("Engine RPM: %s\n", rpm)

	// Send a command to read the engine coolant temperature
	temp := obd.SendOBDCommand("0105", port)
	fmt.Printf("Engine Coolant Temperature: %s°C\n", temp)

	// Send a command to retrieve Diagnostic Trouble Codes (DTCs)
	dtcResponse := obd.SendOBDCommand("03", port)
	dtcs := obd.ParseDTCs(dtcResponse)
	fmt.Printf("Diagnostic Trouble Codes: %v\n", dtcs)
}
