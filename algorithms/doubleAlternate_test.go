package algorithms

import (
	"testing"

	m "github.com/ntindall/moduluschecking/models"
	"github.com/stretchr/testify/assert"
)

func TestDoubleAlternate(t *testing.T) {
	b := m.CreateBankAccount("499273", "12345678")
	sc := m.SortCodeData{
		Algorithm:      "FOO",
		Weights:        []int{2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1},
		ExceptionValue: -1,
		Next:           nil,
		LineNumber:     42,
	}

	assert.Equal(t, 0, DoubleAlternate(b, sc, 0), "Default case in the manual")
	assert.Equal(t, 2, DoubleAlternate(b, sc, 2), "Take into account the initial sum with a remainder")
	assert.Equal(t, 0, DoubleAlternate(b, sc, 10), "Take into account the initial sum without a remainder")
}
