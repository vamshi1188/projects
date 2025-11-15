package functions

import (
	"fmt"
	"os"
)

func DeleteAllNotes() {
	err := os.Remove("notes.txt")

	if err != nil {
		fmt.Println("error in deleting notes", err)
		return
	}
	fmt.Println("Deleted all notes")
}
