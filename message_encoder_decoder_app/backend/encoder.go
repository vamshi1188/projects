package main

import "fmt"

// EncodeMessage converts a string into a sequence of 3-digit ASCII numbers
func EncodeMessage(msg string) string {
	encoded := ""
	for _, ch := range msg {
		encoded += fmt.Sprintf("%03d", int(ch)) // pad each char to 3 digits
	}
	return encoded
}
