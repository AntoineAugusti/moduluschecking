package main

import (
	"fmt"

	"github.com/AntoineAugusti/modulus-checking/models"
	"github.com/AntoineAugusti/modulus-checking/parsers"
	"github.com/AntoineAugusti/modulus-checking/resolvers"
)

func main() {
	parser := parsers.CreateFileParser()
	resolver := resolvers.NewResolver(parser)

	bankAccount := models.CreateBankAccount("089999", "66374958")
	fmt.Println(resolver.IsValid(bankAccount))
}
