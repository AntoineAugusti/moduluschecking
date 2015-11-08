package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortCodeSlice(t *testing.T) {
	b := CreateBankAccount("123456", "12345678")

	expected := []int{1, 2, 3, 4, 5, 6}

	assert.Equal(t, expected, b.SortCodeSlice())
}

func TestAccountNumberSlice(t *testing.T) {
	b := CreateBankAccount("123456", "12345678")

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	assert.Equal(t, expected, b.AccountNumberSlice())
}

func TestMergeAccountDetails(t *testing.T) {
	b := CreateBankAccount("123456", "12345678")
	expected := []int{
		1, 2, 3, 4, 5, 6,
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	assert.Equal(t, expected, b.MergeAccountDetails())
}

func TestNumberAtPosition(t *testing.T) {
	b := CreateBankAccount("123456", "12345678")

	// Test sort code
	assert.Equal(t, 1, b.NumberAtPosition("u"))
	assert.Equal(t, 2, b.NumberAtPosition("v"))
	assert.Equal(t, 3, b.NumberAtPosition("w"))
	assert.Equal(t, 4, b.NumberAtPosition("x"))
	assert.Equal(t, 5, b.NumberAtPosition("y"))
	assert.Equal(t, 6, b.NumberAtPosition("z"))

	// Test bank account
	assert.Equal(t, 1, b.NumberAtPosition("a"))
	assert.Equal(t, 2, b.NumberAtPosition("b"))
	assert.Equal(t, 3, b.NumberAtPosition("c"))
	assert.Equal(t, 4, b.NumberAtPosition("d"))
	assert.Equal(t, 5, b.NumberAtPosition("e"))
	assert.Equal(t, 6, b.NumberAtPosition("f"))
	assert.Equal(t, 7, b.NumberAtPosition("g"))
	assert.Equal(t, 8, b.NumberAtPosition("h"))
}

func TestCreateBankAccount(t *testing.T) {
	b := CreateBankAccount("123456", "12345678")

	assert.Equal(t, "123456", b.SortCode)
	assert.Equal(t, "12345678", b.AccountNumber)

	// Panic cases
	assert.Panics(t, func() { CreateBankAccount("123456", "12345678901") }, "Should panic for account number of 11 digits")
	assert.Panics(t, func() { CreateBankAccount("123456", "12345") }, "Should panic for account number of 5 digits")
	assert.Panics(t, func() { CreateBankAccount("12345", "1234567890") }, "Should panic for sort code of less than 6 digits")

	// 10 digits account number
	coOp := CreateBankAccount("089283", "1234567890")
	assert.Equal(t, "089283", coOp.SortCode, "Sort code from Co-Operative Bank PLC")
	assert.Equal(t, "12345678", coOp.AccountNumber, "Keep only the first 8 digits for an account number from the Co-Operative Bank PLC")

	nationalWestminsterDashes := CreateBankAccount("123456", "01-23456789")
	assert.Equal(t, "123456", nationalWestminsterDashes.SortCode, "Sort code from National Westminster Bank PLC")
	assert.Equal(t, "23456789", nationalWestminsterDashes.AccountNumber, "Keep only the last 8 digits and remove dashes for an account number from National Westminster Bank PLC")

	nationalWestminster := CreateBankAccount("123456", "0123456789")
	assert.Equal(t, "123456", nationalWestminster.SortCode, "Sort code from National Westminster Bank PLC")
	assert.Equal(t, "23456789", nationalWestminster.AccountNumber, "Keep only the last 8 digits and remove dashes for an account number from National Westminster Bank PLC")

	// 9 digits account number
	nineDigits := CreateBankAccount("123456", "123456789")
	assert.Equal(t, "123451", nineDigits.SortCode, "Last digit of sort code is first digit of account number")
	assert.Equal(t, "23456789", nineDigits.AccountNumber, "Account number keeps the last 8 digits")

	// 7 digits account number
	sevenDigits := CreateBankAccount("123456", "1234567")
	assert.Equal(t, "01234567", sevenDigits.AccountNumber, "Zero is added in first position")

	// 6 digits
	sixDigits := CreateBankAccount("123456", "123456")
	assert.Equal(t, "00123456", sixDigits.AccountNumber, "2 zeros are added at the beginning")
}
