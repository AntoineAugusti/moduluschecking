package helpers

import (
	"testing"

	m "github.com/AntoineAugusti/modulus-checking/models"
	"github.com/stretchr/testify/assert"
)

func TestMergeBankAccountDetails(t *testing.T) {
	b := m.CreateBankAccount("123456", "12345678")
	expected := []int{
		1, 2, 3, 4, 5, 6,
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	assert.Equal(t, expected, MergeBankAccountDetails(b))
}

func TestAddDigits(t *testing.T) {
	assert.Equal(t, 3, AddDigits(12))
	assert.Equal(t, 18, AddDigits(99))
	assert.Equal(t, 1, AddDigits(1))
}

func TestLetterToNumber(t *testing.T) {
	b := m.CreateBankAccount("123456", "12345678")

	// Test sort code
	assert.Equal(t, 1, LetterToNumber(b, "u"))
	assert.Equal(t, 2, LetterToNumber(b, "v"))
	assert.Equal(t, 3, LetterToNumber(b, "w"))
	assert.Equal(t, 4, LetterToNumber(b, "x"))
	assert.Equal(t, 5, LetterToNumber(b, "y"))
	assert.Equal(t, 6, LetterToNumber(b, "z"))

	// Test bank account
	assert.Equal(t, 1, LetterToNumber(b, "a"))
	assert.Equal(t, 2, LetterToNumber(b, "b"))
	assert.Equal(t, 3, LetterToNumber(b, "c"))
	assert.Equal(t, 4, LetterToNumber(b, "d"))
	assert.Equal(t, 5, LetterToNumber(b, "e"))
	assert.Equal(t, 6, LetterToNumber(b, "f"))
	assert.Equal(t, 7, LetterToNumber(b, "g"))
	assert.Equal(t, 8, LetterToNumber(b, "h"))
}

func TestAddLeadingZeros(t *testing.T) {
	assert.Equal(t, "000012", AddLeadingZeros(12))
	assert.Equal(t, "000000", AddLeadingZeros(0))
	assert.Equal(t, "123456", AddLeadingZeros(123456))
}

func TestToInt(t *testing.T) {
	assert.Equal(t, 1, ToInt("1"))
}

func TestStringSliceToIntSlice(t *testing.T) {
	expected := []int{42, 1337}
	assert.Equal(t, expected, StringSliceToIntSlice([]string{"42", "1337"}))
}
