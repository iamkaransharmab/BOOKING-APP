package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference" //in golang if you declare the var , then you need use it  ELES will get error
	//fmt.Println(conferenceName)

	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conconferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName) //%T - placeholder for data type

	fmt.Printf("Welcome to %v booking application\n", conferenceName)                                             //for format output wiht placedholder read - pkg.go.dev/fmt                                             //\n - new line & %v placeholder for formatted ouput and we do subsituues reference value with it                   //println - for new line and fmt is pkg top print output in Golang
	fmt.Printf("We have total of %v tickets and %v  are still available.\n", conferenceTickets, remainingTickets) //println - for new line and
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user for their name

	userName = "Keera"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets) //%v or fmt placeholder works with fmt.Printf  -IMP*

}
