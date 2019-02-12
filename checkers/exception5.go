package checkers

import (
	m "github.com/ntindall/moduluschecking/models"
)

type Exception5Checker struct {
	Substitutions map[string]string
}

// Determine if the checker is able to validate the bank account
func (e Exception5Checker) Handles(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	return sc.IsException(5)
}

// Tell if the bank account is valid
func (e Exception5Checker) IsValid(b m.BankAccount, sc m.SortCodeData, attempt int) bool {
	if !e.Handles(b, sc, attempt) {
		panic("Should be exception of type 5")
	}

	if substitution, hasKey := e.Substitutions[b.SortCode]; hasKey {
		b.SortCode = substitution
	}

	// First attempt
	if attempt == 1 {
		checkDigit := b.NumberAtPosition("g")
		remainder := GeneralChecker{}.RemainderFromRegularCheck(b, sc)
		if remainder == 0 && checkDigit == 0 {
			return true
		}
		if remainder == 1 {
			return false
		}

		return (11 - remainder) == checkDigit
	}

	// Second attempt
	checkDigit := b.NumberAtPosition("h")
	remainder := GeneralChecker{}.RemainderFromRegularCheck(b, sc)
	if remainder == 0 && checkDigit == 0 {
		return true
	}

	return (10 - remainder) == checkDigit
}
