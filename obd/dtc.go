package obd

import (
	"fmt"
)

// ParseDTCs parses the response from the "03" OBD-II command and returns a list of DTCs
func ParseDTCs(response string) []string {
	// Clean up the response
	response = response[:len(response)-1] // Remove trailing '>'

	dtcCount := (len(response) - 2) / 4
	dtcs := make([]string, 0, dtcCount)

	for i := 2; i < len(response); i += 4 {
		dtc := response[i : i+4]
		if len(dtc) == 4 {
			dtcs = append(dtcs, decodeDTC(dtc))
		}
	}

	return dtcs
}

// decodeDTC converts a raw DTC string to a human-readable format
func decodeDTC(dtc string) string {
	if len(dtc) != 4 {
		return ""
	}

	// The first character
	char1 := string(dtc[0])
	switch char1 {
	case "0", "1", "2", "3":
		char1 = "P"
	case "4", "5", "6", "7":
		char1 = "C"
	case "8", "9", "A", "B":
		char1 = "B"
	case "C", "D", "E", "F":
		char1 = "U"
	}

	return fmt.Sprintf("%s%s", char1, dtc[1:])
}
