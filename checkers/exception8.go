package checkers

import (
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Perform the check for the exception 8
func PerformException8Check(b m.BankAccount, scData m.SortCodeData) bool {
	if !scData.IsException(8) {
		panic("Should be exception of type 8")
	}

	b.SortCode = "090126"
	return PerformRegularCheck(b, scData)
}
