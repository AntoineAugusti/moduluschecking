package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBankAccount(t *testing.T) {
	b := CreateBankAccount("123456", "12345678")

	assert.Equal(t, "123456", b.SortCode)
	assert.Equal(t, "12345678", b.AccountNumber)
}

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
