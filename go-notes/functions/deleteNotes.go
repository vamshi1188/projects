package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DeleteNotes() {

	var word string

	fmt.Scan(&word)

	DeleteLinesContaining(word)

}

func DeleteLinesContaining(word string) {
	inputFile := "notes.txt"
	tempFile := "temp.txt"

	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer in.Close()

	out, err := os.Create(tempFile)
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer out.Close()

	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	for scanner.Scan() {
		line := scanner.Text()
		// If line does NOT contain the word, keep it
		if !strings.Contains(line, word) {
			fmt.Fprintln(writer, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	writer.Flush()
	in.Close()
	out.Close()

	// Replace old file with updated one
	err = os.Rename(tempFile, inputFile)
	if err != nil {
		fmt.Println("Error replacing file:", err)
		return
	}

	fmt.Println("✅ Lines containing", word, "were deleted successfully!")
}
