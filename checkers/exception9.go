package checkers

import (
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Perform the check for the exception 9
func PerformException9Check(b m.BankAccount, scData m.SortCodeData, weights map[string]m.SortCodeData) bool {
	if !scData.IsException(9) {
		panic("Should be exception of type 9")
	}

	scData.Weights = WeightsForException2Or9(b, scData)
	if PerformRegularCheck(b, scData) {
		return true
	}

	// Try to replace the sort code
	b.SortCode = "309634"
	return PerformRegularCheck(b, weights[b.SortCode])
}
