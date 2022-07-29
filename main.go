package main

import "fmt"

func main() {
	conferenceName := "p2well's Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint =  50
	var bookings [50]string
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	
	var firstName string
	var lastName string
	var eMail string
	var userTickets uint

	// ask user for their name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your e-mail address: ")
	fmt.Scan(&eMail)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array size: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v\n", firstName, lastName, userTickets, eMail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
}
