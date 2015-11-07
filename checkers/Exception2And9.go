package checkers

import (
	"github.com/AntoineAugusti/modulus-checking/helpers"
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Get weights to use for a bank account following the exception 2 or 9.
func WeightsForException2Or9(b m.BankAccount, sc m.SortCodeData) (weights []int) {
	if !(sc.IsException(2) || sc.IsException(9)) {
		panic("Expected exception 2 or exception 9 sort code")
	}

	a := helpers.LetterToNumber(b, "a")
	g := helpers.LetterToNumber(b, "g")

	// Default weights
	weights = sc.Weights

	switch {
	case a != 0 && g != 9:
		weights = []int{
			0, 0, 1, 2, 5, 3,
			6, 4, 8, 7, 10, 9, 3, 1,
		}
	case a != 0 && g == 9:
		weights = []int{
			0, 0, 0, 0, 0, 0,
			0, 0, 8, 7, 10, 9, 3, 1,
		}
	}

	return
}
