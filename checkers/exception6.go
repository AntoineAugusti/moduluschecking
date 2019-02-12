package checkers

import (
	m "github.com/ntindall/moduluschecking/models"
)

type Exception6Checker struct {
}

// Determine if the checker is able to validate the bank account
func (e Exception6Checker) Handles(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return sc.IsException(6)
}

// Tell if the bank account is valid
func (e Exception6Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 6")
	}

	if e.isForeignCurrency(b) {
		return true
	}

	return GeneralChecker{}.IsValid(b, sc, attempt)
}

// Check if a bank account is in a foreign currency
func (e Exception6Checker) isForeignCurrency(b m.BankAccount) bool {
	// if a = 4, 5, 6, 7 or 8, and g and h are the same,
	// the accounts are for a foreign currency and the checks cannot be used
	a := b.NumberAtPosition("a")
	g := b.NumberAtPosition("g")
	h := b.NumberAtPosition("h")

	return a >= 4 && a <= 8 && (g == h)
}
