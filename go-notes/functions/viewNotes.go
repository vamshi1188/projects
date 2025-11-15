package functions

import (
	"fmt"
	"os"
)

func ViewNotes() {
	fmt.Println("-- view notes --")

	file, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := os.ReadFile("notes.txt")
	if err != nil {
		fmt.Println("error in opening file ", err)
	}
	fmt.Println(string(data))
}
