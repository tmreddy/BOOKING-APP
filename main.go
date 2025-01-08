package main

import (
	"fmt"
	"strings"
)

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


	for {
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


		isValidName := len(firstName) >=2 && len(lasttName) >= 0
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets
		 

		if isValidEmail && isValidName && isValidTickets {
			remainingTickets = remainingTickets -  userTickets

			bookings = append(bookings,  firstName + " " + lasttName) 


			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lasttName,  userTickets, email )

			fmt.Printf("%v tickets remaining, for %v \n",remainingTickets,conferenceName)

			firstNames := []string{}

			for  _ , booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			
			}

			fmt.Printf("The firstname of the bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets have been booked. Come back next year")
				break
			}	
		} else {
			if !isValidName {
				fmt.Printf("Your name is invalid, try again\n")
			}
			if !isValidEmail {
				fmt.Printf("Your email is invalid, try again\n")
			}
			if !isValidTickets {
				fmt.Printf("Your tickets is invalid, try again\n")
			}
		}
	}
}
