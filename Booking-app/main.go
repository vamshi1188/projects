package main

import (
	"fmt"
	"strconv"
	"strings"
)

var conferenceName = " go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUSers()

	for {
		firstName, lastName, email, userTickets := getUSerInput()

		isValidname, isValidemail, isValidticketnumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidemail && isValidname && isValidticketnumber {
			bookTicket(userTickets, firstName, lastName, email)
			// call first name function
			firstNames := getFirstnames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end the programe
				fmt.Println("our conference tickets are sold out .come back next year")
				break
			}

		} else {
			if !isValidname {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidemail {
				fmt.Println("email address you entered is invalid")
			}
			if !isValidticketnumber {
				fmt.Println(" number of tickets you entered is invalid")
			}

		}

	}

}

func greetUSers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")
}

func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidname := len(firstName) >= 2 && len(lastName) >= 2
	isValidemail := strings.Contains(email, "@") && strings.Contains(email, ".com") || strings.Contains(email, ".in")
	isValidticketnumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidname, isValidemail, isValidticketnumber
}
func getUSerInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//map
	var userDAta = make(map[string]string)
	userDAta["firstName"] = firstName
	userDAta["lastName"] = lastName
	userDAta["email"] = email
	userDAta["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userDAta)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("thank you %v %v for booking %v tickets You will receive confermation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}
