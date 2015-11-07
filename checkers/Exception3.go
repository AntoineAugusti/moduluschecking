package checkers

import (
	"github.com/AntoineAugusti/modulus-checking/helpers"
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Check if a bank account matches the criteria of the exception 3
func Exception3(account m.BankAccount) bool {
	// If c=6 or c=9
	c := helpers.LetterToNumber(account, "c")

	return c == 6 || c == 9
}
