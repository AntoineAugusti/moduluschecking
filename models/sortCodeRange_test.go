package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortCodeRangeHasException(t *testing.T) {
	sc := SortCodeRange{
		Start:          1,
		End:            10,
		Algorithm:      "foo",
		Weights:        []int{},
		ExceptionValue: 0,
		LineNumber:     42,
	}

	assert.False(t, sc.HasException())

	sc.ExceptionValue = 1

	assert.True(t, sc.HasException())
}
