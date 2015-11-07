package helpers

import (
	"strconv"

	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Merge the sort code and the account number of a
// bank account into a single slice
func MergeBankAccountDetails(b m.BankAccount) []int {
	return append(b.SortCodeSlice(), b.AccountNumberSlice()...)
}

// Perform an addition over the digits
// of a 2 digits number
func AddDigits(nb int) int {
	if nb > 99 || nb < 0 {
		panic("Number should be between 0 and 99")
	}
	return nb%10 + nb/10
}

// Get the integer value from a letter, according to the defined code:
// Letters between u and z select a digit from the sort code
// Letters between a and h select a digit from the account number
func LetterToNumber(b m.BankAccount, letter string) int {
	nb := MergeBankAccountDetails(b)
	switch {
	case letter == "u":
		return nb[0]
	case letter == "v":
		return nb[1]
	case letter == "w":
		return nb[2]
	case letter == "x":
		return nb[3]
	case letter == "y":
		return nb[4]
	case letter == "z":
		return nb[5]
	case letter == "a":
		return nb[6]
	case letter == "b":
		return nb[7]
	case letter == "c":
		return nb[8]
	case letter == "d":
		return nb[9]
	case letter == "e":
		return nb[10]
	case letter == "f":
		return nb[11]
	case letter == "g":
		return nb[12]
	case letter == "h":
		return nb[13]
	}

	panic("Unknow letter")
}

// Add leading zeros to a number to have a string
// of length 6
func AddLeadingZeros(nb int) string {
	s := strconv.Itoa(nb)
	// Make sure the string has got a length of 6
	for len(s) < 6 {
		s = "0" + s
	}
	return s
}

// Try to transform a string to an int
// and panic if it fails
func ToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

// Convert a slice of strings to a slice of integers
func StringSliceToIntSlice(slice []string) (res []int) {
	for _, element := range slice {
		res = append(res, ToInt(element))
	}

	return res
}
