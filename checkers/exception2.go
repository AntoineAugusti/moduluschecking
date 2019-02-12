package checkers

import (
	m "github.com/ntindall/moduluschecking/models"
)

type Exception2Checker struct {
}

// Determine if the checker is able to validate the bank account
func (e Exception2Checker) Handles(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return sc.IsException(2)
}

// Tell if the bank account is valid
func (e Exception2Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 2")
	}

	sc.Weights = WeightsForException2Or9(b, sc)

	return GeneralChecker{}.IsValid(b, sc, attempt)
}
