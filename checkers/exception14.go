package checkers

import (
	m "github.com/ntindall/moduluschecking/models"
)

type Exception14Checker struct {
}

// Determine if the checker is able to validate the bank account
func (e Exception14Checker) Handles(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return sc.IsException(14) && attempt == 2
}

// Tell if the bank account is valid
func (e Exception14Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 14 at attempt 2")
	}

	h := b.NumberAtPosition("h")
	if h >= 2 && h <= 8 {
		return false
	}
	// Remove the 1st digit from the accout number and insert a 0
	// as the 1st digit for check purposes
	b.AccountNumber = "0" + b.AccountNumber[0:len(b.AccountNumber)-1]

	return GeneralChecker{}.IsValid(b, sc, attempt)
}
