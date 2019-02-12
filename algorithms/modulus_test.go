package algorithms

import (
	"testing"

	m "github.com/AntoineAugusti/moduluschecking/models"
	"github.com/stretchr/testify/assert"
)

func TestModulus(t *testing.T) {
	b := m.CreateBankAccount("000000", "58177632")
	sc := m.SortCodeData{
		Algorithm:      "FOO",
		Weights:        []int{0, 0, 0, 0, 0, 0, 7, 5, 8, 3, 4, 6, 2, 1},
		ExceptionValue: -1,
		Next:           nil,
		LineNumber:     42,
	}

	assert.Equal(t, 0, Modulus(b, 11, sc), "Default case in the manual")
	assert.Equal(t, 6, Modulus(b, 10, sc), "Take into account another modulo - 10")
	assert.Equal(t, 0, Modulus(b, 176, sc), "Take into account another modulo - 176")
}
