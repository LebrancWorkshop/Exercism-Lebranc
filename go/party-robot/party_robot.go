package partyrobot

import (
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!";  
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"; 
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcomeMsg := "Welcome to my party, " + name + "!";
	
	var tableNumber string; 

	if table >= 0 && table <= 9 {
		tableNumber = "00" + strconv.Itoa(table);	
	} else if table >= 10 && table <= 99 {
		tableNumber = "0" + strconv.Itoa(table);
	} else {
		tableNumber = strconv.Itoa(table);  
	}

	assignMsg := "You have been assigned to table " + tableNumber + ". Your table is " + direction + ", exactly " + strconv.FormatFloat(distance, 'f', 1, 64) + " meters from here.";  
	seatMateMsg := "You will be sitting next to " + neighbor + ".";
	return welcomeMsg + "\n" + assignMsg + "\n" + seatMateMsg;   
}
