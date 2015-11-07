package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasException(t *testing.T) {
	sc := base()

	assert.False(t, sc.HasException())

	sc.ExceptionValue = 1
	assert.True(t, sc.HasException())
}

func TestHasNext(t *testing.T) {
	sc := base()

	assert.False(t, sc.HasNext())

	other := base()
	sc.Next = &other
	assert.True(t, sc.HasNext())
}

func TestIsException(t *testing.T) {
	sc := base()
	sc.ExceptionValue = 3

	assert.False(t, sc.IsException(4))
	assert.True(t, sc.IsException(3))
}

func TestFollowsExceptions(t *testing.T) {
	sc := base()
	sc.ExceptionValue = 3

	assert.False(t, sc.FollowsExceptions(3, 4))

	other := base()
	other.ExceptionValue = 4
	sc.Next = &other
	assert.True(t, sc.FollowsExceptions(3, 4))
	assert.False(t, sc.FollowsExceptions(3, 5))
}

func base() SortCodeData {
	return SortCodeData{
		Algorithm:      "foo",
		Weights:        []int{},
		ExceptionValue: 0,
		Next:           nil,
		LineNumber:     42,
	}
}
