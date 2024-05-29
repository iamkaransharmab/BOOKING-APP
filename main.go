package main

import (
	"fmt"
)

func main() {

	var conferenceName string = "Go Conference" //in golang if you declare the var , then you need use it  ELES will get error // := easy syntax can be called as synthetic sugar
	//fmt.Println(conferenceName)

	const conferenceTickets int = 50
	var remainingTickets uint = 50 //remaining tickets should be positive valuea

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conconferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName) //%T - placeholder for data type

	fmt.Printf("Welcome to %v booking application\n", conferenceName)                                             //for format output wiht placedholder read - pkg.go.dev/fmt                                             //\n - new line & %v placeholder for formatted ouput and we do subsituues reference value with it                   //println - for new line and fmt is pkg top print output in Golang
	fmt.Printf("We have total of %v tickets and %v  are still available.\n", conferenceTickets, remainingTickets) //println - for new line and
	fmt.Println("Get your tickets here to attend")

	//user
	var firstName string
	var lastName string
	var email string
	var userTickets uint //userTickets should be positive valuea
	//var bookings [50]string //array for booking ticket logic
	var bookings []string //slice for booking //var bookings = []string{} or bookings := []string{}
	//ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your booked ticktes:")
	fmt.Scan(&userTickets)

	//Book ticket logic
	remainingTickets = remainingTickets - userTickets

	//bookings[0] = firstName + " " + lastName //refer booking array
	bookings = append(bookings, firstName+" "+lastName) //in slice we don't ned index , it just append the next value as its dynamic array //adding value is differ form array in slice but getting or retrieving a value from slide is similar

	//this is clear as commnet for now -VideoTime: 1:10
	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("slice type: %T\n", bookings)
	// fmt.Printf("slice length: %v\n", len(bookings))

	// fmt.Scan(&firstName)    //user input funtion //&variable name -it helps to get input and so user can get desired outpur by giving input
	//user input funtion //&-pointer help to get input

	//RUF
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets) //pointer

	// userName = "Keera"
	// userTickets = 2
	// fmt.Printf("User  %v  %v booked %v tickets and his Email addres %v.\n", firstName, lastName, userTickets, email) //%v or fmt placeholder works with fmt.Printf  -IMP*

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email) //%v or fmt placeholder works withme, us email) //%v or fmt placeholder works with
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)                                                                //%v or fmt placeholder worksiningTickets, con) //%v or fmt placeholder works
	fmt.Printf("These are all bookings: %v\n", bookings)

}
