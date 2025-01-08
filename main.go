package main

import (
	"fmt"
	"strconv"
)

const conferenceTickets = 50
var conferenceName = "Go Conference" // camel casing
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	

	greetUsers()

	for {
		 
		firstName, lasttName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := validateUserInputs(firstName, lasttName, email, userTickets)

		if isValidEmail && isValidName && isValidTickets {
			
			bookTickets(firstName, lasttName, email, userTickets)

			firstNames := getFirstName()

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

func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")

}

func getFirstName() []string {

	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])

	}
	return firstNames
}


func getUserInput() (string, string, string, uint){
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

	return firstName, lasttName, email, userTickets
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = map[string]string{}

	userData["firstName"] = firstName
	userData["lastName"] = lastName 
	userData["email"] = email
	userData["tickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining, for %v \n", remainingTickets, conferenceName)
}