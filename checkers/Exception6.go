package checkers

import (
	"github.com/AntoineAugusti/moduluschecking/helpers"
	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Perform the check for the exception 6
func PerformException6Check(b m.BankAccount, scData m.SortCodeData) bool {
	if !scData.IsException(6) {
		panic("Should be exception of type 6")
	}

	if isForeignCurrency(b) {
		return true
	}

	return PerformRegularCheck(b, scData)
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
