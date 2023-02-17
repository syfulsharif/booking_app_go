package main

import "fmt"

func main() {
	var conferenceName = "Go Conferenece"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and currently %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}
