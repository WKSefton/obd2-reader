package obd

import (
	"log"
	"time"

	"golang.org/x/sys/windows"
)

// InitializeELM327 sends initialization commands to the ELM327 device
func InitializeELM327(port windows.Handle) {
	writeToPort(port, "ATZ\r")
	time.Sleep(2 * time.Second)
	readResponse(port)

	writeToPort(port, "ATE0\r")
	time.Sleep(500 * time.Millisecond)
	readResponse(port)

	writeToPort(port, "ATSP0\r")
	time.Sleep(500 * time.Millisecond)
	readResponse(port)
}

// SendOBDCommand sends an OBD-II command and returns the response
func SendOBDCommand(command string, port windows.Handle) string {
	writeToPort(port, command+"\r")
	time.Sleep(500 * time.Millisecond)
	return readResponse(port)
}

// Helper functions for writing to and reading from the serial port
func writeToPort(port windows.Handle, data string) {
	buf := []byte(data)
	var written uint32
	err := windows.WriteFile(port, buf, &written, nil)
	if err != nil {
		log.Fatalf("Failed to write to serial port: %v", err)
	}
}

func readResponse(port windows.Handle) string {
	buf := make([]byte, 1024)
	var read uint32
	err := windows.ReadFile(port, buf, &read, nil)
	if err != nil {
		log.Fatalf("Failed to read from serial port: %v", err)
	}
	response := string(buf[:read])
	return response
}
