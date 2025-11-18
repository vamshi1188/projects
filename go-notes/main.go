package main

import (
	"fmt"
	"go-notes/functions"
)

func main() {
	fmt.Println("Welcome to go-notes📒")
	DisplayOpt()
	var UserInput int
	fmt.Scan(&UserInput)
	switch UserInput {
	case 1:
		functions.AddNotes()
	case 2:
		functions.ViewNotes()
	case 3:
		functions.DeleteNotes()
	case 4:
		functions.EditNotes()
	case 5:
		functions.DeleteAllNotes()

	}
}

func DisplayOpt() {
	fmt.Println("choose options below")
	fmt.Println("1 - Add notes")
	fmt.Println("2 - View notes")
	fmt.Println("3 - Delete notes")
	fmt.Println("4 - Edit notes")
	fmt.Println("5 - Delete All notes")

}
