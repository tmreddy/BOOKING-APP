package main

import "fmt"

func main(){
	var conferenceName = "Go Conference" // camel casing
	const conferenceTickets = 50
	var remainingTickets uint= 50
	var bookings []string

	fmt.Printf("confirenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferenceTickets,remainingTickets,conferenceName)

	//fmt.Println("Welcome to our", conferenceName,  "booking application")
	fmt.Printf("Welcome to our %v booking application\n",conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string  
	var lasttName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter you last name: ")
	fmt.Scan(&lasttName)
	
	fmt.Println("Enter you email: ")
	fmt.Scan(&email)
	
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets -  userTickets
	
	bookings = append(bookings,  firstName + " " + lasttName) 


	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lasttName,  userTickets, email )

	fmt.Printf("%v tickets remaining, for %v \n",remainingTickets,conferenceName)

	fmt.Printf("These are all the bookings: %v\n", bookings)

}
