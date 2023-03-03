package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conferenece"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and currently %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// var bookings = [50]string{}
	for {
		var bookings []string

		//Assigning Value to the slice using Index
		// bookings[0] = "Sharif"
		// bookings[1] = "Nupur"

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
		bookings = append(bookings, firstName+" "+lastName)

		// userName = "Tom"
		// userTickets = 2
		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first element of slice:  %v\n", bookings[0])
		// fmt.Printf("The type of slice: %T\n", bookings)
		// fmt.Printf("The length of slice: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets, your tickets have been sent to %v. \n", firstName, lastName, userTickets, email)
		fmt.Printf("There are %v tickets left for '%v' \n", remainingTickets, conferenceName)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are first names of all our bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			//end the program
			fmt.Printf("Our %v is booked out, try again later", conferenceName)
			break
		}
	}

}
