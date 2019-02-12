package algorithms

import (
	"github.com/ntindall/moduluschecking/helpers"
	m "github.com/ntindall/moduluschecking/models"
)

// Perform the double alternate algorithm and return
// the remainder of the operation
func DoubleAlternate(b m.BankAccount, data m.SortCodeData, sum int) (remainder int) {
	numbers := b.MergeAccountDetails()

	weights := data.Weights
	for i, nb := range numbers {
		sum += helpers.AddDigits(weights[i] * nb)
	}

	remainder = sum % 10

	return
}
