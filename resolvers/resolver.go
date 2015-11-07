package resolvers

import (
	"github.com/AntoineAugusti/modulus-checking/checkers"
	"github.com/AntoineAugusti/modulus-checking/helpers"
	m "github.com/AntoineAugusti/modulus-checking/models"
)

type Resolver struct {
	weights       map[string]m.SortCodeData
	substitutions map[string]string
}

// Check if a bank account number is valid
func (r Resolver) IsValid(b m.BankAccount) bool {
	scData := r.weights[b.SortCode]
	secondCheck := true

	// First check
	firstCheck := r.checkForSortCodeData(b, scData, 1)

	if checkers.Exception3(b) {
		return firstCheck
	}

	if scData.FollowsExceptions(2, 9) {
		if firstCheck {
			return true
		}
	}

	if scData.IsException(14) {
		if firstCheck {
			return true
		}
		return r.checkForSortCodeData(b, scData, 2)
	}

	// Second check if required
	if scData.HasNext() {
		secondCheck = r.checkForSortCodeData(b, *scData.Next, 2)
	}

	if scData.FollowsExceptions(2, 9) || scData.FollowsExceptions(10, 11) || scData.FollowsExceptions(12, 13) {
		return firstCheck || secondCheck
	}

	return firstCheck && secondCheck
}

func (r Resolver) checkForSortCodeData(b m.BankAccount, scData m.SortCodeData, attempt int) bool {
	if scData.HasException() {
		switch {
		case scData.IsException(1):
			return checkers.DoubleAlternate(b, scData, 27) == 0
		case scData.IsException(2):
			scData.Weights = checkers.WeightsForException2Or9(b, scData)
			return performRegularCheck(b, scData)
		case scData.IsException(4):
			g := helpers.LetterToNumber(b, "g")
			h := helpers.LetterToNumber(b, "h")
			checkDigit := g*10 + h
			return remainderFromRegularCheck(b, scData) == checkDigit
		case scData.IsException(5):
			if substitution, hasKey := r.substitutions[b.SortCode]; hasKey {
				b.SortCode = substitution
			}
			if attempt == 1 {
				checkDigit := helpers.LetterToNumber(b, "g")
				remainder := remainderFromRegularCheck(b, scData)
				if remainder == 0 && checkDigit == 0 {
					return true
				}
				if remainder == 1 {
					return false
				}
				return (11 - remainder) == checkDigit
			} else {
				checkDigit := helpers.LetterToNumber(b, "h")
				remainder := remainderFromRegularCheck(b, scData)
				if remainder == 0 && checkDigit == 0 {
					return true
				}
				return (10 - remainder) == checkDigit
			}
		case scData.IsException(6):
			if isForeignCurrency := checkers.Exception6(b); isForeignCurrency {
				return true
			}
		case scData.IsException(7):
			if checkers.Exception7(b) {
				zeros := []int{0, 0, 0, 0, 0, 0, 0, 0}
				scData.Weights = append(zeros, scData.Weights[8:]...)
				return performRegularCheck(b, scData)
			}
		case scData.IsException(8):
			b.SortCode = "090126"
			return performRegularCheck(b, scData)
		case scData.IsException(9):
			scData.Weights = checkers.WeightsForException2Or9(b, scData)
			if performRegularCheck(b, scData) {
				return true
			}
			// Try to replace the sort code
			b.SortCode = "309634"
			res := performRegularCheck(b, r.weights[b.SortCode])
			return res
		case scData.IsException(10):
			if checkers.Exception10(b) {
				zeros := []int{0, 0, 0, 0, 0, 0, 0, 0}
				scData.Weights = append(zeros, scData.Weights[8:]...)
				return performRegularCheck(b, scData)
			}
		case scData.IsException(14):
			if attempt == 2 {
				h := helpers.LetterToNumber(b, "h")
				if h >= 2 && h <= 8 {
					return false
				}
				b.AccountNumber = "0" + b.AccountNumber[0:len(b.AccountNumber)-1]
				return performRegularCheck(b, scData)
			}
		}
	}

	return performRegularCheck(b, scData)
}

func performRegularCheck(b m.BankAccount, scData m.SortCodeData) bool {
	return remainderFromRegularCheck(b, scData) == 0
}

func remainderFromRegularCheck(b m.BankAccount, scData m.SortCodeData) int {
	switch {
	case scData.Algorithm == "DBLAL":
		return checkers.DoubleAlternate(b, scData, 0)
	case scData.Algorithm == "MOD11":
		remainder := checkers.Modulus(b, 11, scData)
		return remainder
	case scData.Algorithm == "MOD10":
		remainder := checkers.Modulus(b, 10, scData)
		return remainder
	}

	panic("Should have algo")
}

// Construct a new Resolver and automatically read the
// initialization files
func NewResolver(parser m.Parser) Resolver {
	resolver := Resolver{
		weights:       parser.Weights(),
		substitutions: parser.Substitutions(),
	}

	return resolver
}
