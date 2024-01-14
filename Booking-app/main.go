package main

import "fmt"

func main() {
	var conferenceName = " go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user to enter their name

	userName = "tom"
	userTickets = 2
	fmt.Println(userName)
	fmt.Printf("user %v booked %v tickets\n", userName, userTickets)

}
