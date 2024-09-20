package techpalace

import "strings"

func repeatStars(qtd int) string {
	return strings.Repeat("*", qtd)
}

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return repeatStars(numStarsPerLine) + "\n" + welcomeMsg + "\n" + repeatStars(numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	oldMsg = strings.ReplaceAll(oldMsg, "\n", "")
	return strings.TrimSpace(oldMsg)
}
