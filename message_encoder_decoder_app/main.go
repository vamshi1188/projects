package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message to encode: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	encoded := EncodeMessage(message)
	fmt.Println("Encoded:", encoded)

	decoded := DecodeMessage(encoded)
	fmt.Println("Decoded:", decoded)
}
