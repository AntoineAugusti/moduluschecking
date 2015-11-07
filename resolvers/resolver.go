package resolvers

import (
	"github.com/AntoineAugusti/moduluschecking/checkers"
	m "github.com/AntoineAugusti/moduluschecking/models"
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

	if checkers.FollowsException3(b, scData) {
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

	// Some sort code require only one of the 2 checks to be successful
	if scData.FollowsExceptions(2, 9) || scData.FollowsExceptions(10, 11) || scData.FollowsExceptions(12, 13) {
		return firstCheck || secondCheck
	}

	return firstCheck && secondCheck
}

// Perform the check for a bank account and a given attempt.
func (r Resolver) checkForSortCodeData(b m.BankAccount, scData m.SortCodeData, attempt int) bool {
	if scData.HasException() {
		switch {
		case scData.IsException(1):
			return checkers.PerformException1Check(b, scData)
		case scData.IsException(2):
			return checkers.PerformException2Check(b, scData)
		case scData.IsException(4):
			return checkers.PerformException4Check(b, scData)
		case scData.IsException(5):
			return checkers.PerformException5Check(b, scData, r.substitutions, attempt)
		case scData.IsException(6):
			return checkers.PerformException6Check(b, scData)
		case scData.IsException(7):
			return checkers.PerformException7Check(b, scData)
		case scData.IsException(8):
			return checkers.PerformException8Check(b, scData)
		case scData.IsException(9):
			return checkers.PerformException9Check(b, scData, r.weights)
		case scData.IsException(10):
			return checkers.PerformException10Check(b, scData)
		case scData.IsException(14):
			return checkers.PerformException14Check(b, scData, attempt)
		}
	}

	return checkers.PerformRegularCheck(b, scData)
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
