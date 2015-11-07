package checkers

import (
	"github.com/AntoineAugusti/moduluschecking/helpers"
	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Perform the check for the exception 4
func PerformException4Check(b m.BankAccount, scData m.SortCodeData) bool {
	if !scData.IsException(4) {
		panic("Should be exception of type 4")
	}

	g := helpers.LetterToNumber(b, "g")
	h := helpers.LetterToNumber(b, "h")
	checkDigit := g*10 + h

	return RemainderFromRegularCheck(b, scData) == checkDigit
}
