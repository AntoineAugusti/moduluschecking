package checkers

import (
	m "github.com/ntindall/moduluschecking/models"
)

type Exception10Checker struct {
}

// Determine if the checker is able to validate the bank account
func (e Exception10Checker) Handles(account m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !sc.IsException(10) {
		return false
	}

	// if ab = 09 or ab = 99 and g = 9
	a := account.NumberAtPosition("a")
	b := account.NumberAtPosition("b")
	g := account.NumberAtPosition("g")

	return (a == 0 || a == 9) && b == 9 && g == 9
}

// Tell if the bank account is valid
func (e Exception10Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 10")
	}

	// Put 8 zeros at the beginning of the weights
	// and keep the original weights after
	zeros := []int{0, 0, 0, 0, 0, 0, 0, 0}
	sc.Weights = append(zeros, sc.Weights[8:]...)

	return GeneralChecker{}.IsValid(b, sc, attempt)
}
