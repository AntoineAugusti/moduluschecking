package models

import (
	"strconv"
)

// Represents a UK bank account
type BankAccount struct {
	SortCode      string
	AccountNumber string
}

// The sort code has an integers slice
func (b BankAccount) SortCodeSlice() []int {
	return toSlice(b.SortCode)
}

// The account number has an integers slice
func (b BankAccount) AccountNumberSlice() []int {
	return toSlice(b.AccountNumber)
}

// Create a BankAccount structure from a sort code and an account number
func CreateBankAccount(sortCode, accountNumber string) BankAccount {
	return BankAccount{
		SortCode:      sortCode,
		AccountNumber: accountNumber,
	}
}

// Create integers slice from an integers slice.
func toSlice(str string) []int {
	var res []int

	for _, c := range str {
		v, _ := strconv.Atoi(string(c))
		res = append(res, v)
	}

	return res
}
