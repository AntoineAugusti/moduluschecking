package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDigits(t *testing.T) {
	assert.Equal(t, 3, AddDigits(12))
	assert.Equal(t, 18, AddDigits(99))
	assert.Equal(t, 1, AddDigits(1))
}

func TestAddLeadingZerosToNumber(t *testing.T) {
	assert.Equal(t, "000012", AddLeadingZerosToNumber(12))
	assert.Equal(t, "000000", AddLeadingZerosToNumber(0))
	assert.Equal(t, "123456", AddLeadingZerosToNumber(123456))
}

func TestAddLeadingZeros(t *testing.T) {
	assert.Equal(t, "000ab", AddLeadingZeros("ab", 5))
	assert.Equal(t, "ab", AddLeadingZeros("ab", 2))
}

func TestToInt(t *testing.T) {
	assert.Equal(t, 1, ToInt("1"))
}

func TestStringToIntSlice(t *testing.T) {
	expected := []int{1, 3, 3, 7}

	assert.Equal(t, expected, StringToIntSlice("1337"))
	assert.Nil(t, StringToIntSlice(""))
}

func TestStringSliceToIntSlice(t *testing.T) {
	expected := []int{42, 1337}
	assert.Equal(t, expected, StringSliceToIntSlice([]string{"42", "1337"}))
}

func TestRemoveDashes(t *testing.T) {
	assert.Equal(t, "123", RemoveDashes("1-23"))
	assert.Equal(t, "123", RemoveDashes("123"))
	assert.Equal(t, "123", RemoveDashes("1-2-3-"))
	assert.Equal(t, "", RemoveDashes(""))
}
