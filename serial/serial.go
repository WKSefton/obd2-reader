package serial

import (
	"log"

	"golang.org/x/sys/windows"
)

// OpenSerialPort opens the serial port for communication
func OpenSerialPort(portName string) (windows.Handle, error) {
	serialPort := `\\.\` + portName
	port, err := windows.CreateFile(
		windows.StringToUTF16Ptr(serialPort),
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		0, nil, windows.OPEN_EXISTING, 0, 0,
	)
	if err != nil {
		return windows.InvalidHandle, err
	}
	configureSerialPort(port)
	return port, nil
}

// CloseSerialPort closes the serial port
func CloseSerialPort(port windows.Handle) {
	err := windows.CloseHandle(port)
	if err != nil {
		log.Fatalf("Failed to close serial port: %v", err)
	}
}
