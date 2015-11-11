package models

// A checker is in charge of validating a bank account number
type Checker interface {
	// Determine if the checker is able to validate the bank account
	Handles(b BankAccount, sc SortCodeData, attempt int) bool
	// Tell if the bank account is valid
	IsValid(b BankAccount, sc SortCodeData, attempt int) bool
}

type Checkers []Checker
