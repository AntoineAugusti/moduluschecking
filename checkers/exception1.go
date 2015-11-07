package checkers

import (
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Perform the check for the exception 1
func PerformException1Check(b m.BankAccount, scData m.SortCodeData) bool {
	if !scData.IsException(1) {
		panic("Should be exception of type 1")
	}

	return DoubleAlternate(b, scData, 27) == 0
}
