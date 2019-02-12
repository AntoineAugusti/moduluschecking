package checkers

import (
	m "github.com/ntindall/moduluschecking/models"
)

type Exception9Checker struct {
	Weights map[string]m.SortCodeData
}

// Determine if the checker is able to validate the bank account
func (e Exception9Checker) Handles(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return sc.IsException(9)
}

// Tell if the bank account is valid
func (e Exception9Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 9")
	}

	sc.Weights = WeightsForException2Or9(b, sc)
	if (GeneralChecker{}.IsValid(b, sc, attempt)) {
		return true
	}

	// Try to replace the sort code
	b.SortCode = "309634"

	return GeneralChecker{}.IsValid(b, e.Weights[b.SortCode], attempt)
}
