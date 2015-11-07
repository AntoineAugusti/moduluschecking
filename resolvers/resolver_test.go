package resolvers

import (
	"testing"

	m "github.com/AntoineAugusti/moduluschecking/models"
	"github.com/AntoineAugusti/moduluschecking/parsers"
	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	parser := parsers.CreateFileParser()
	res := NewResolver(parser)

	assert.True(t, res.IsValid(m.CreateBankAccount("089999", "66374958")), "1- Pass modulus 10 check")
	assert.True(t, res.IsValid(m.CreateBankAccount("107999", "88837491")), "2- Pass modulus 11 check")
	assert.True(t, res.IsValid(m.CreateBankAccount("202959", "63748472")), "3- Pass modulus 11 check and double alternate checks")
	assert.True(t, res.IsValid(m.CreateBankAccount("871427", "46238510")), "4- Exception 10 and 11 where first check passes and second check fails")
	assert.True(t, res.IsValid(m.CreateBankAccount("872427", "46238510")), "5- Exception 10 and 11 where first check fails and second check passes")
	assert.True(t, res.IsValid(m.CreateBankAccount("871427", "09123496")), "6- Exception 10 where in the account number ab=09 and the g=9. The first check passes and the second check fails.")
	assert.True(t, res.IsValid(m.CreateBankAccount("871427", "99123496")), "7- Exception 10 where in the account number ab=99 and the g=9. The first check passes and the second check fails.")
	assert.True(t, res.IsValid(m.CreateBankAccount("820000", "73688637")), "8- Exception 3, and the sorting code is the start of a range. As c=6 the second check should be ignored.")
	assert.True(t, res.IsValid(m.CreateBankAccount("827999", "73988638")), "9- Exception 3, and the sorting code is the end of a range. As c=9 the second check should be ignored.")
	assert.True(t, res.IsValid(m.CreateBankAccount("827101", "28748352")), "10- Exception 3. As c <> 6 or 9 perform both checks pass.")
	assert.True(t, res.IsValid(m.CreateBankAccount("134020", "63849203")), "11- Exception 4 where the remainder is equal to the checkdigit.")
	assert.True(t, res.IsValid(m.CreateBankAccount("118765", "64371389")), "12- Exception 1 - ensures that 27 has been added to the accumulated total and passes double alternate modulus check.")
	assert.True(t, res.IsValid(m.CreateBankAccount("200915", "41011166")), "13- Exception 6 where the account fails standard check but is a foreign currency account.")
	assert.True(t, res.IsValid(m.CreateBankAccount("938611", "07806039")), "14- Exception 5 where the check passes.")
	assert.True(t, res.IsValid(m.CreateBankAccount("938600", "42368003")), "15- Exception 5 where the check passes with substitution.")
	assert.True(t, res.IsValid(m.CreateBankAccount("938063", "55065200")), "16- Exception 5 where both checks produce a remainder of 0 and pass.")
	assert.True(t, res.IsValid(m.CreateBankAccount("772798", "99345694")), "17- Exception 7 where passes but would fail the standard check.")
	assert.True(t, res.IsValid(m.CreateBankAccount("086090", "06774744")), "18- Exception 8 where the check passes.")
	assert.True(t, res.IsValid(m.CreateBankAccount("309070", "02355688")), "19- Exception 2 & 9 where the first check passes.")
	assert.True(t, res.IsValid(m.CreateBankAccount("309070", "12345668")), "20- Exception 2 & 9 where the first check fails and second check passes with substitution.")
	assert.True(t, res.IsValid(m.CreateBankAccount("309070", "12345677")), "21- Exception 2 & 9 where a <> 0 and g <> 9 and passes.")
	assert.True(t, res.IsValid(m.CreateBankAccount("309070", "99345694")), "22- Exception 2 & 9 where a <> 0 and g = 9 and passes.")
	assert.False(t, res.IsValid(m.CreateBankAccount("938063", "15764273")), "23- Exception 5 where the first checkdigit is correct and the second incorrect.")
	assert.False(t, res.IsValid(m.CreateBankAccount("938063", "15764264")), "24- Exception 5 where the first checkdigit is incorrect and the second correct.")
	assert.False(t, res.IsValid(m.CreateBankAccount("938063", "15763217")), "25- Exception 5 where the first checkdigit is incorrect with a remainder of 1.")
	assert.False(t, res.IsValid(m.CreateBankAccount("118765", "64371388")), "26- Exception 1 where it fails double alternate check.")
	assert.False(t, res.IsValid(m.CreateBankAccount("203099", "66831036")), "27- Pass modulus 11 check and fail double alternate check.")
	assert.False(t, res.IsValid(m.CreateBankAccount("203099", "58716970")), "28- Fail modulus 11 check and pass double alternate check.")
	assert.False(t, res.IsValid(m.CreateBankAccount("089999", "66374959")), "29- Fail modulus 10 check.")
	assert.False(t, res.IsValid(m.CreateBankAccount("107999", "88837493")), "30- Fail modulus 11 check.")
	assert.True(t, res.IsValid(m.CreateBankAccount("074456", "12345112")), "31- Exception 12/13 where passes modulus 11 check.")
	assert.True(t, res.IsValid(m.CreateBankAccount("070116", "34012583")), "32- Exception 12/13 where passes modulus 11 check.")
	assert.True(t, res.IsValid(m.CreateBankAccount("074456", "11104102")), "33- Exception 12/13 where fails the modulus 11 check, but passes the modulus 10 check.")
	assert.True(t, res.IsValid(m.CreateBankAccount("180002", "00000190")), "34- Exception 14 where the first check fails and the second check passes.")
	assert.True(t, res.IsValid(m.CreateBankAccount("180002", "98093517")), "Additional- Exception 14 where the first check passes.")
}
