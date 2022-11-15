package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customerUpper := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, "	+ customerUpper 
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starBorder := strings.Repeat("*", numStarsPerLine)
	return starBorder + "\n" + welcomeMsg + "\n" + starBorder  
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.Replace(oldMsg, "*", "", len(oldMsg))
	newMsgNoBlank := strings.Replace(newMsg, "\n", "", len(newMsg))
	finalMsg := strings.TrimSpace(newMsgNoBlank)  
	return finalMsg  
}
