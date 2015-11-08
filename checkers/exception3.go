package checkers

import (
	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Check if we follow the criteria of the exception 3
func FollowsException3(b m.BankAccount, scData m.SortCodeData) bool {
	c := b.NumberAtPosition("c")

	hasNextAnd3 := scData.HasNext() && scData.Next.IsException(3)

	return (scData.IsException(3) || hasNextAnd3) && (c == 6 || c == 9)
}
