package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conferenece"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	greetUser(conferenceName, conferenceTickets, remainingTickets)

	var bookings []string
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets, your tickets have been sent to %v. \n", firstName, lastName, userTickets, email)
			fmt.Printf("There are %v tickets left for '%v' \n", remainingTickets, conferenceName)
			// Print First Names
			printFirstNames(bookings)

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

func greetUser(confName string, confTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application.\n", confName)
	fmt.Printf("We have total of %v tickets and currently %v tickets left.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("These are first names of all our bookings are: %v\n", firstNames)
}
