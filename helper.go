package main 

import "strings"

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool){

	isValidName := len(firstName) >= 2 && len(lastName) >= 0
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}
