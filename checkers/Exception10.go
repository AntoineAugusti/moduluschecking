package checkers

import (
	"github.com/AntoineAugusti/modulus-checking/helpers"
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Check if a bank account matches the criteria of the exception 10
func Exception10(account m.BankAccount) bool {
	// if ab = 09 or ab = 99 and g = 9
	a := helpers.LetterToNumber(account, "a")
	b := helpers.LetterToNumber(account, "b")
	g := helpers.LetterToNumber(account, "g")

	return (a == 0 || a == 9) && b == 9 && g == 9
}
