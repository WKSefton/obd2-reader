package serial

import (
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

// COMMTIMEOUTS defines the timeout settings for the serial port
type COMMTIMEOUTS struct {
	ReadIntervalTimeout         uint32
	ReadTotalTimeoutMultiplier  uint32
	ReadTotalTimeoutConstant    uint32
	WriteTotalTimeoutMultiplier uint32
	WriteTotalTimeoutConstant   uint32
}

func configureSerialPort(port windows.Handle) {
	var dcb windows.DCB
	dcb.DCBlength = uint32(unsafe.Sizeof(dcb))

	// Get the current serial port settings
	if err := windows.GetCommState(port, &dcb); err != nil {
		log.Fatalf("Failed to get serial port state: %v", err)
	}

	// Configure DCB settings
	dcb.BaudRate = 9600                      // Set baud rate to 9600
	dcb.ByteSize = 8                         // 8 data bits
	dcb.Parity = windows.NOPARITY            // No parity bit
	dcb.StopBits = windows.ONESTOPBIT        // 1 stop bit
	dcb.Flags = dcbFlagsEnableDTR(dcb.Flags) // Enable DTR (Data Terminal Ready)

	// Set the updated settings
	if err := windows.SetCommState(port, &dcb); err != nil {
		log.Fatalf("Failed to set serial port state: %v", err)
	}

	// Set timeouts
	timeouts := COMMTIMEOUTS{
		ReadIntervalTimeout:         50,
		ReadTotalTimeoutMultiplier:  10,
		ReadTotalTimeoutConstant:    50,
		WriteTotalTimeoutMultiplier: 10,
		WriteTotalTimeoutConstant:   50,
	}
	if err := windows.SetCommTimeouts(port, (*windows.CommTimeouts)(unsafe.Pointer(&timeouts))); err != nil {
		log.Fatalf("Failed to set serial port timeouts: %v", err)
	}
}

// dcbFlagsEnableDTR enables the DTR (Data Terminal Ready) control flow
func dcbFlagsEnableDTR(flags uint32) uint32 {
	const DTR_CONTROL_ENABLE uint32 = 0x01
	return flags | DTR_CONTROL_ENABLE
}
