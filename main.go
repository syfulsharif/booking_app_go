package main

import "fmt"

func main() {
	var conferenceName = "Go Conferenece"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", conferenceName, "booking application.")
	fmt.Println("We have total of", conferenceTickets, "tickets and currently", remainingTickets, "tickets left.")
	fmt.Println("Get your tickets here to attend.")

}
