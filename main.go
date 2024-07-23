package main

import "fmt"
/*
	You can create a variable and Go will dynamically assign a type using this
		example conferenceName := "Ruby Conf Africa 2024" for strings
				ticketsAvailable := 150 for integers
	Arrays in go are fixed
		example var bookings = [50]string
	Slices are better since they are dynamic arrays
		empty slice => bookings := []string{} OR var bookings []string{}

	Loops in Go
		

*/
func main() {
	var remainingTickets uint = 50
	conferenceName := "Ruby Conf Africa"

	var bookings []string
	// empty slice => bookings := []string{} OR var bookings []string{}

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName) // the ampasant sign is called a pointer. Used to infer the memory location where the variable is stored
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter your number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
	fmt.Printf("There are all our bookings: %v\n", bookings)
}
