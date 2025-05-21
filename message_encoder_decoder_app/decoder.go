package main

import (
	"strconv"
)

// DecodeMessage converts a sequence of 3-digit ASCII numbers back into the original string
func DecodeMessage(code string) string {
	decoded := ""
	for i := 0; i < len(code); i += 3 {
		part := code[i : i+3]
		num, err := strconv.Atoi(part)
		if err == nil {
			decoded += string(rune(num))
		}
	}
	return decoded
}
