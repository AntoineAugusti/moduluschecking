package algorithms

import (
	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Perform the modulus algorithm with a given modulus
// and return the remainder of the operation
func Modulus(b m.BankAccount, modulus int, data m.SortCodeData) (remainder int) {
	numbers := b.MergeAccountDetails()
	sum := 0
	weight := data.Weights

	for i, nb := range numbers {
		sum += weight[i] * nb
	}

	remainder = sum % modulus

	return
}
