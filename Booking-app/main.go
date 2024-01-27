package main

import "fmt"

func main() {
	conferenceName := " go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask the  user to enter their name
	fmt.Print("Enter ur first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter ur last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter ur email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("thank you %v %v for booking %v tickets. You will receive confermation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}
