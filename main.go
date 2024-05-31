package main

import (
	"fmt"
	"strings" //this pkg for strings.Fields()
)

func main() {

	var conferenceName string = "Go Conference" //in golang if you declare the var , then you need use it  ELES will get error // := easy syntax can be called as synthetic sugar
	//fmt.Println(conferenceName)

	const conferenceTickets int = 50
	var remainingTickets uint = 50 //remaining tickets should be positive valuea

	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conconferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName) //%T - placeholder for data type

	//putting this greating code inside the greet function
	//fmt.Printf("Welcome to %v booking application\n", conferenceName)        //for format output wiht placedholder read - pkg.go.dev/fmt                                             //\n - new line & %v placeholder for formatted ouput and we do subsituues reference value with it                   //println - for new line and fmt is pkg top print output in Golang
	//fmt.Printf("We have total of %v tickets and %v  are still available.\n", conferenceTickets, remainingTickets) //println - for new line and
	//fmt.Println("Get your tickets here to attend")

	greetUser(conferenceName, conferenceTickets, remainingTickets) //we need always pass all the arguemenst that written inside the decaler func

	//Loops
	for { //we cna defince conditon for- fro loop als,that how much time or long it should be run ,.like for tru-it indicate that loop is always teue anf its infinite loop and also can set another condiotns like remainingTickets && userTickets <= 50

		//video time 2:13 hrs - Now will create function for user Input and will call same like validUserInpu function and grretUser and getFirstNames in main function.
		//now here, we will collect fucntion with values and store in varibles

		//user
		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint //userTickets should be positive valuea
		// //var bookings [50]string //array for booking ticket logic
		// var bookings []string //slice for booking //var bookings = []string{} or bookings := []string{}
		// //ask user for their name
		// fmt.Println("Enter your first name:")
		// fmt.Scan(&firstName)

		// fmt.Println("Enter your Last name:")
		// fmt.Scan(&lastName)

		// fmt.Println("Enter your Email address:")
		// fmt.Scan(&email)

		// fmt.Println("Enter your booked ticktes:")
		// fmt.Scan(&userTickets)

		//decalring its as function :
		// isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// isValidEmail := strings.Contains(email, "@") //contains for check email and it will give outpt bool value //its form strings pkg
		// isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets //for validTicketNumber.that shoudl be psotive
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets) //this is declare for isValidateUserInput function

		if isValidName && isValidEmail && isValidTicketNumber { ////userTickets <= remainingTickets -this commend from below if condition
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

			//we are creating fucntion for print firstNames
			firstNames := getFirstNames(bookings) //decaled firstNames as variable ,so can call it directly by main
			fmt.Printf("The first Names of bookings are: %v\n", firstNames)

			// //we want just output with user's firstNames
			// firstNames := []string{}           //slice deacalring
			// for _, booking := range bookings { //for array & slices -range provide the index and value for each-element
			// 	//in Golang, Underscore(_) used to indetify unused variables ,as act as Black Identifier
			// 	var names = strings.Fields(booking) //strings.Fields()- spilts the string with space as separator
			// 	firstNames = append(firstNames, names[0])
			// }
			// fmt.Printf("The first Names of bookings are: %v\n", firstNames)

			//now we r gooona implement if-else statement, So we can break for loop,so if remainingTickets go out 0 then it will stop the program wiht if-else condition

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out, Come back next year.")
				break
			} //now the htingsis when user booking gretaher than avaiable or rmeaning tikctes, the confercne dhad not ended and its givingg randome ref value,so wee need ot fix that by applying conditons  userTcikers > remainingtickets
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

			fmt.Println("We apologize, but your ticket request is invalid. Please try again.") //now we are saying invalid but we can optimise the userInput as invalid and  printhe same input in Output
		}

		// fmt.Printf("We only have %v tickets remaining , so you cna't book %v tickets\n", remainingTickets, userTickets)
		//break// But as if now we want to tell customer to book the ticket  again under the avaibl conference ticket
		//continue //by continue keyword .customer will able to book ticket again uder tha availability of conference tickets

	}

}

//creating function with fucn keyword

func greetUser(confName string, confTickets int, remainingTickets uint) { //function decalre but it can only be excuted after its called,inside the main fucntion
	//var confName string , so confName is variable that I have declared inside the function
	fmt.Printf("Welcome to our %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v  are still available.\n", confTickets, remainingTickets) //println - for new line and
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames(bookings []string) []string {
	//we want just output with user's firstNames
	firstNames := []string{}           //slice deacalring
	for _, booking := range bookings { //for array & slices -range provide the index and value for each-element
		//in Golang, Underscore(_) used to indetify unused variables ,as act as Black Identifier
		var names = strings.Fields(booking) //strings.Fields()- spilts the string with space as separator
		firstNames = append(firstNames, names[0])
	}
	//fmt.Printf("The first Names of bookings are: %v\n", firstNames)
	return firstNames //if we want print firstNames in the Main function ,then we need to sue returnkeyword and also defien the type of return function ,in the beginning
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) { //(bool,boo,bool) written in parenthesis cos od multiple data types values & now need to collect them inside the mainfcuntion ,there the isValidateuserInput func declared andstore in the varibale
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")                              //contains for check email and it will give outpt bool value //its form strings pkg
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets //for validTicketNumber.that shoudl be psotive

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint //userTickets should be positive valuea
	//var bookings [50]string //array for booking ticket logic
	//var bookings []string //slice for booking //var bookings = []string{} or bookings := []string{}
	//ask user for their name
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
