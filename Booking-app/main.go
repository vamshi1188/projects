package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := " go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask the  user to enter their name
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidname := len(firstName) >= 2 && len(lastName) >= 2
		isValidemail := strings.Contains(email, "@")
		isValidticketnumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidemail && isValidname && isValidticketnumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("thank you %v %v for booking %v tickets. You will receive confermation mail at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end the programe
				fmt.Println("our conference tickets are sold out .come back next year")
				break
			}

		} else {
			fmt.Println("your input data is invalid , try again ")

		}

	}

}
