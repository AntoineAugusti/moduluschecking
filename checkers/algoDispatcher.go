package checkers

import (
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Get the remainder of the check
func RemainderFromRegularCheck(b m.BankAccount, scData m.SortCodeData) int {
	switch {
	case scData.Algorithm == "DBLAL":
		return DoubleAlternate(b, scData, 0)
	case scData.Algorithm == "MOD11":
		return Modulus(b, 11, scData)
	case scData.Algorithm == "MOD10":
		return Modulus(b, 10, scData)
	}

	panic("Should have algo")
}

// Check that the remainder of a regular check is zero
func PerformRegularCheck(b m.BankAccount, scData m.SortCodeData) bool {
	return RemainderFromRegularCheck(b, scData) == 0
}
