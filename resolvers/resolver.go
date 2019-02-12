package resolvers

import (
	"github.com/ntindall/moduluschecking/checkers"
	m "github.com/ntindall/moduluschecking/models"
)

type Resolver struct {
	// Associate a sort code to its data
	weights map[string]m.SortCodeData
	// Sort code substitutions map
	substitutions map[string]string
	// Associate an exception ID to a specific checker
	exceptionCheckers map[int]m.Checker
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
	// If the sort code follows an exception
	if scData.HasException() {
		// Try to find a checker for this specific exception
		if checkerFound, hasKey := r.exceptionCheckers[scData.ExceptionValue]; hasKey {
			// Check that this checker can actually handle this case
			if checkerFound.Handles(b, scData, attempt) {
				// Validate the bank account number
				return checkerFound.IsValid(b, scData, attempt)
			}
		}
	}

	// General case
	return checkers.GeneralChecker{}.IsValid(b, scData, attempt)
}

// Construct a new Resolver and automatically read the
// initialization files
func NewResolver(parser m.Parser) Resolver {
	weights, substitutions := parser.Weights(), parser.Substitutions()

	resolver := Resolver{
		weights:       weights,
		substitutions: substitutions,
		exceptionCheckers: map[int]m.Checker{
			1:  checkers.Exception1Checker{},
			2:  checkers.Exception2Checker{},
			4:  checkers.Exception4Checker{},
			5:  checkers.Exception5Checker{Substitutions: substitutions},
			6:  checkers.Exception6Checker{},
			7:  checkers.Exception7Checker{},
			8:  checkers.Exception8Checker{},
			9:  checkers.Exception9Checker{Weights: weights},
			10: checkers.Exception10Checker{},
			14: checkers.Exception14Checker{},
		},
	}

	return resolver
}
