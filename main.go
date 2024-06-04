package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var RemainingTickets uint = 50
var bookings = make([]userData, 0) //list of maps ,with 0 size //VideoTime 2:58Hrs - changing list of map to list of  userData Ty[pe] {/make([]map[string]string, 0)}

type userData struct { //type keyword create a new type with nae, you specify for anydata type ,In STruct can defiend mix-Data types //struct is lighetweight class ,which doesn't support Inheritance
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, RemainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email) //for concurrecy justy and go keyword in starting

		firstNames := getFirstNames()
		fmt.Printf("The first Names of bookings are: %v\n", firstNames)

		if RemainingTickets == 0 {

			fmt.Println("Our conference is booked out, Come back next year.")

		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't seems correct")
		}
		if !isValidTicketNumber {
			fmt.Println("ticket number you entered is invalid")
		}

		fmt.Println("We apologize, but your ticket request is invalid. Please try again.")
	}

	wg.Wait()

}

func greetUser() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v  are still available.\n", conferenceTickets, RemainingTickets) //println - for new line and
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)

		firstNames = append(firstNames, booking.firstName) //with STRUCT, it predefined the values ,and asigne dwiht struct.values
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your booked ticktes:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	RemainingTickets = RemainingTickets - userTickets

	//create a map for user

	var userData = userData{ // this will userData objects , sowe won't need that userData conversiion(line 111-115)
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	} // make(map[string]string)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email

	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData) //userData = firstName + " " + lastName
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", RemainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //Sprintf - to prin strings
	fmt.Println("############")
	fmt.Printf("Sending ticket :\n %v \n to email address %v\n", ticket, email)
	fmt.Println("###########")

	wg.Done()
}
