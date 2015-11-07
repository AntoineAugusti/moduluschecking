package checkers

import (
	"github.com/AntoineAugusti/modulus-checking/helpers"
	m "github.com/AntoineAugusti/modulus-checking/models"
)

// Perform the check for the exception 14
func PerformException14Check(b m.BankAccount, scData m.SortCodeData, attempt int) bool {
	if !scData.IsException(14) {
		panic("Should be exception of type 14")
	}

	if attempt == 2 {
		h := helpers.LetterToNumber(b, "h")
		if h >= 2 && h <= 8 {
			return false
		}
		b.AccountNumber = "0" + b.AccountNumber[0:len(b.AccountNumber)-1]
		return PerformRegularCheck(b, scData)
	}

	return PerformRegularCheck(b, scData)
}
