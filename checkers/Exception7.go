package checkers

import (
	"github.com/AntoineAugusti/modulus-checking/helpers"
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Check if a bank account matches the criteria of the exception 7
func Exception7(account m.BankAccount) bool {
	// If g=9
	g := helpers.LetterToNumber(account, "g")

	return g == 9
}
