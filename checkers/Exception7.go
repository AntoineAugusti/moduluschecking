package checkers

import (
	"github.com/AntoineAugusti/modulus-checking/helpers"
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Perform the check for the exception 7
func PerformException7Check(b m.BankAccount, scData m.SortCodeData) bool {
	if !scData.IsException(7) {
		panic("Should be exception of type 7")
	}

	if isException7(b) {
		zeros := []int{0, 0, 0, 0, 0, 0, 0, 0}
		scData.Weights = append(zeros, scData.Weights[8:]...)
		return PerformRegularCheck(b, scData)
	}

	return PerformRegularCheck(b, scData)
}

// Check if a bank account matches the criteria of the exception 7
func isException7(account m.BankAccount) bool {
	// If g=9
	g := helpers.LetterToNumber(account, "g")

	return g == 9
}
