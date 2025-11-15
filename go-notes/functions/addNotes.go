package functions

import (
	"bufio"
	"fmt"
	"os"
)

func AddNotes() {
	fmt.Println("-- Add notes --")
	fmt.Println("Enter notes in this order \nTitle:note")

	// input from user
	reader := bufio.NewReader(os.Stdin)

	Notes, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {

		fmt.Println("error in opening file", err)
	}
	defer file.Close()

	_, err = file.WriteString(Notes)

	if err != nil {
		fmt.Println("error in writing file", err)
		return
	}

}
