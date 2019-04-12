package checkers

import (
	"github.com/AntoineAugusti/moduluschecking/algorithms"
	m "github.com/AntoineAugusti/moduluschecking/models"
)

type GeneralChecker struct {
}

// Determine if the checker is able to validate the bank account
func (e GeneralChecker) Handles(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return true
}

// Tell if the bank account is valid
func (e GeneralChecker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return e.RemainderFromRegularCheck(b, sc) == 0
}

// RemainderFromRegularCheck gets the remainder of the check
func (e GeneralChecker) RemainderFromRegularCheck(b m.BankAccount, scData m.SortCodeData) int {
	switch {
	case scData.Algorithm == "DBLAL":
		return algorithms.DoubleAlternate(b, scData, 0)
	case scData.Algorithm == "MOD11":
		return algorithms.Modulus(b, 11, scData)
	case scData.Algorithm == "MOD10":
		return algorithms.Modulus(b, 10, scData)
	}

	panic("Should have algo")
}
