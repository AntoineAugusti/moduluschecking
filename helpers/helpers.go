package helpers

import (
	"strconv"
	"strings"
)

// Perform an addition over the digits
// of a 2 digits number
func AddDigits(nb int) int {
	if nb > 99 || nb < 0 {
		panic("Number should be between 0 and 99")
	}
	return nb%10 + nb/10
}

// Add leading zeros to a number to have a string
// of length 6
func AddLeadingZerosToNumber(nb int) string {
	return AddLeadingZeros(strconv.Itoa(nb), 6)
}

// Add leading zeros to a string to have a string
// of a desired length
func AddLeadingZeros(s string, desiredLength int) string {
	for len(s) < desiredLength {
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

// Convert a string containing multiple digits
// to a string of integers
func StringToIntSlice(s string) []int {
	res := []string{}
	for _, c := range s {
		res = append(res, string(c))
	}
	return StringSliceToIntSlice(res)
}

// Convert a slice of strings to a slice of integers
func StringSliceToIntSlice(slice []string) (res []int) {
	for _, element := range slice {
		res = append(res, ToInt(element))
	}

	return res
}

// Remove dashes from a string
func RemoveDashes(s string) string {
	return strings.Replace(s, "-", "", -1)
}
