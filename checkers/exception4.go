package checkers

import (
	m "github.com/ntindall/moduluschecking/models"
)

type Exception4Checker struct {
}

// Determine if the checker is able to validate the bank account
func (e Exception4Checker) Handles(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return sc.IsException(4)
}

// Tell if the bank account is valid
func (e Exception4Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 4")
	}

	g := b.NumberAtPosition("g")
	h := b.NumberAtPosition("h")
	checkDigit := g*10 + h

	return GeneralChecker{}.RemainderFromRegularCheck(b, sc) == checkDigit
}
