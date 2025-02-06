package main

import (
	"fmt"
	"strconv"
	 "time"
)

type customer struct{
	FirstName string
	LastName string
	Email string
	UserTickets int
	Attended bool
	AttendedTickets int
}

func main(){
	// Create a map
	// map[keyType]valueType
	// keyType can be any type that can be compared with ==, != operators
	// valueType can be any type

	bookings := make([]map[string]string,0)
	customers := make([]customer,0) 

	var firstName, lastName, email string 
	var tickets int

	// for {

		fmt.Println("Enter first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter email ")
		fmt.Scan(&email)
		fmt.Println("Enter tickets ")
		fmt.Scan(&tickets)

		// userData["fName"] = firstName
		// userData["lName"] = lastName
		// userData["emailAddress"] = email
		// userData["UserTickets"] = tickets
		// printMap(userData)

		// userData[firstName] = firstName
		// userData[lastName] = lastName
		// userData[email] = email
		// userData[tickets] = tickets

		bookings = append(bookings, createMap(firstName, lastName, email, tickets))
		customers = append(customers, createCustomer(firstName, lastName, email, tickets))

		printMap(bookings)
		printCustomer(customers)
		go sendEmail(customers)
	// }
}

func createCustomer(fName string, lName string, userEmail string, userTickets int) customer{
	
	customer := customer{
		FirstName: fName,
		LastName: lName,
		Email: userEmail,
		UserTickets: userTickets,
	}

	return customer
}

func createMap(fName string, lName string, userEmail string, userTickets int) map[string]string{
	
	myMap := make(map[string]string)

	myMap["FirstName"] =fName
	myMap["LastName"] = lName
	myMap["Email"] = userEmail
	myMap["UserTickets"] = strconv.Itoa(userTickets)

	return myMap
	
}

func printMap(myMap []map[string]string){
	for key, value := range myMap{
		fmt.Println(key, value)

		// for k,v := range value{
		// 	fmt.Println(k, v)
		// 	fmt.Printf("%v, %T \n",k, v)
		// }
	}
}

func printCustomer(customers []customer){
	for key, value := range customers{
		fmt.Println(key, value)
		fmt.Println(value.FirstName, value.LastName, value.Email, value.UserTickets)
	}
}

func sendEmail(customer []customer){
	time.Sleep(5 * time.Second)
	for _, customer := range customer{
		// sub := "Thank you for booking tickets"
		// body := "Thank you for booking tickets. Your tickets are confirmed"
		// from := "booking@bookmyshow.com"
		to := customer.Email 
		fmt.Println("Sending email to ", to)
	}
}

func markAttendence(customer customer, attended bool, noOfAttendes int){
	customer.Attended = attended
	customer.AttendedTickets = noOfAttendes
}

// firstName naga
// lastName lakshmi
// email naga@gmail.com
// tickets 2

// firstName naga
// lastName lakshmi
// email naga@gmail.com
// tickets 2
// naga naga
// lakshmi lakshmi
// naga@gmail.com naga@gmail.com
// 2 2

// firstName Mallikarjun
// lastName reddy
// email reddy@gmail.com
// tickets 5
// naga naga
// lakshmi lakshmi
// naga@gmail.com naga@gmail.com
// 2 2
// Mallikarjun Mallikarjun
// reddy reddy
// reddy@gmail.com reddy@gmail.com
// 5 5

// firstName Dines
// naga naga
// reddy reddy
// 2 2
// 5 5
// email dinesh@gmail.com
// lakshmi lakshmi
// naga@gmail.com naga@gmail.com
// Mallikarjun Mallikarjun
// Dines Dines
// dinesh@gmail.com dinesh@gmail.com
// lastName reddy
// tickets 5
// reddy@gmail.com reddy@gmail.com