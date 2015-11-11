package checkers

import (
	m "github.com/AntoineAugusti/moduluschecking/models"
)

type Exception7Checker struct {
}

// Determine if the checker is able to validate the bank account
func (e Exception7Checker) Handles(account m.BankAccount, sc m.SortCodeData, attempt int) bool {
	// If g=9
	g := account.NumberAtPosition("g")

	return sc.IsException(7) && g == 9
}

// Tell if the bank account is valid
func (e Exception7Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 7")
	}

	// Put 8 zeros at the beginning of the weights
	// and keep the original weights after
	zeros := []int{0, 0, 0, 0, 0, 0, 0, 0}
	sc.Weights = append(zeros, sc.Weights[8:]...)

	return GeneralChecker{}.IsValid(b, sc, attempt)
}
