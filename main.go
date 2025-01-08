package main

import "fmt"

func main(){
	var conferenceName = "Go Conference" // camel casing
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("confirenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferenceTickets,remainingTickets,conferenceName)

	//fmt.Println("Welcome to our", conferenceName,  "booking application")
	fmt.Printf("Welcome to our %v booking application\n",conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lasttName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter you last name: ")
	fmt.Scan(&lasttName)
	fmt.Println("Enter you email: ")
	fmt.Scan(&email)
	fmt.Println(&remainingTickets)

	fmt.Printf("User %v booked %v tickets. \n", firstName, userTickets)

}
