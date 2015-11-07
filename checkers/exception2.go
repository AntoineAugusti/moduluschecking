package checkers

import (
	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Perform the check for the exception 2
func PerformException2Check(b m.BankAccount, scData m.SortCodeData) bool {
	if !scData.IsException(2) {
		panic("Should be exception of type 2")
	}

	scData.Weights = WeightsForException2Or9(b, scData)
	return PerformRegularCheck(b, scData)
}
