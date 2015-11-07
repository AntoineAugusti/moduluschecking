package checkers

import (
	"github.com/AntoineAugusti/modulus-checking/helpers"
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Check if a bank account matches the criteria of the exception 6
func Exception6(b m.BankAccount) bool {
	return isForeignCurrency(b)
}

// Check if a bank account is in a foreign currency
func isForeignCurrency(b m.BankAccount) bool {
	// if a = 4, 5, 6, 7 or 8, and g and h are the same,
	// the accounts are for a foreign currency and the checks cannot be used
	a := helpers.LetterToNumber(b, "a")
	g := helpers.LetterToNumber(b, "g")
	h := helpers.LetterToNumber(b, "h")

	return a >= 4 && a <= 8 && (g == h)
}
