package checkers

import (
	"github.com/AntoineAugusti/moduluschecking/algorithms"
	m "github.com/AntoineAugusti/moduluschecking/models"
)

type Exception1Checker struct {
}

// Determine if the checker is able to validate the bank account
func (e Exception1Checker) Handles(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return sc.IsException(1)
}

// Tell if the bank account is valid
func (e Exception1Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 1")
	}

	return algorithms.DoubleAlternate(b, sc, 27) == 0
}
