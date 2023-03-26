package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conferenece"

var remainingTickets uint = 50

var bookings []string

func main() {

	greetUser()

	for {

		//User Input Section
		firstName, lastName, email, userTickets := getUserInput()
		//User Input Validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//Booking Logic
			bookTicket(userTickets, firstName, lastName, email)
			// Print First Names
			fmt.Printf("These are first names of all our bookings are: %v\n", getFirstNames())

			if remainingTickets == 0 {
				//end the program
				fmt.Printf("Our %v is booked out, try again later", conferenceName)
				break
			}
		} else {
			// fmt.Println("You have entered invalid input")
			if !isValidName {
				fmt.Println("You have entered wrong first name or last name")
			}
			if !isValidEmail {
				fmt.Println("Your email address doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The ticket number you entered is invalid")
			}
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and currently %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	//Returning First Names I/O printing
	return firstNames
	// fmt.Printf("These are first names of all our bookings are: %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) ([]string, uint) {
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets, your tickets have been sent to %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets left for '%v' \n", remainingTickets, conferenceName)
	bookings = append(bookings, firstName+" "+lastName)
	return bookings, remainingTickets
}
