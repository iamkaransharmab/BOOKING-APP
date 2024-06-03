package main

import (
	"strings"
)

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) { //(bool,boo,bool) written in parenthesis cos od multiple data types values & now need to collect them inside the mainfcuntion ,there the isValidateuserInput func declared andstore in the varibale
	isValidName := len(firstName) >= 2 && len(lastName) >= 2                  //removed remainingTicktes form argument ,as it splg
	isValidEmail := strings.Contains(email, "@")                              //contains for check email and it will give outpt bool value //its form strings pkg
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets //for validTicketNumber.that shoudl be psotive

	return isValidName, isValidEmail, isValidTicketNumber
}
