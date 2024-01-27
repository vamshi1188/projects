package main

import "fmt"

func main() {
	conferenceName := " go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")

	var userName string
	var userTickets int
	//ask the  user to enter their name
	fmt.Println("Enter ur first name")
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("user %v booked %v tickets\n", userName, userTickets)

}
