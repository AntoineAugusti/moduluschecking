package checkers

import (
	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Perform the check for the exception 5
func PerformException5Check(b m.BankAccount, scData m.SortCodeData, substitutions map[string]string, attempt int) bool {
	if !scData.IsException(5) {
		panic("Should be exception of type 5")
	}

	if substitution, hasKey := substitutions[b.SortCode]; hasKey {
		b.SortCode = substitution
	}

	// First attempt
	if attempt == 1 {
		checkDigit := b.NumberAtPosition("g")
		remainder := RemainderFromRegularCheck(b, scData)
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
	remainder := RemainderFromRegularCheck(b, scData)
	if remainder == 0 && checkDigit == 0 {
		return true
	}
	return (10 - remainder) == checkDigit
}
