package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conferenece"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and currently %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

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

	remainingTickets = remainingTickets - userTickets

	// userName = "Tom"
	// userTickets = 2
	fmt.Printf("Thank you %v %v for booking %v tickets, your tickets have been sent to %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets left for '%v' \n", remainingTickets, conferenceName)
}
