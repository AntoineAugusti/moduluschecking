package checkers

import (
	m "github.com/ntindall/moduluschecking/models"
)

type Exception8Checker struct {
}

// Determine if the checker is able to validate the bank account
func (e Exception8Checker) Handles(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return sc.IsException(8)
}

// Tell if the bank account is valid
func (e Exception8Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 8")
	}

	b.SortCode = "090126"

	return GeneralChecker{}.IsValid(b, sc, attempt)
}
