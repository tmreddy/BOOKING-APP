package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50
var conferenceName = "Go Conference" // camel casing
var remainingTickets uint = 50
var bookings []string

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
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool){

	isValidName := len(firstName) >= 2 && len(lastName) >= 0
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
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

func bookTickets(firstName string, lasttName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lasttName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lasttName, userTickets, email)

	fmt.Printf("%v tickets remaining, for %v \n", remainingTickets, conferenceName)
}