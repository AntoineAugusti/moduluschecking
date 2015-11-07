package checkers

import (
	"github.com/AntoineAugusti/moduluschecking/helpers"
	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Perform the double alternate algorithm and return
// the remainder of the operation
func DoubleAlternate(b m.BankAccount, data m.SortCodeData, sum int) (remainder int) {
	numbers := helpers.MergeBankAccountDetails(b)

	weight := data.Weights
	for i, nb := range numbers {
		sum += helpers.AddDigits(weight[i] * nb)
	}

	remainder = sum % 10

	return
}
