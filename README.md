[![Travis CI](https://img.shields.io/travis/AntoineAugusti/modulus-checking/master.svg?style=flat-square)](https://travis-ci.org/AntoineAugusti/modulus-checking)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/antoineaugusti/modulus-checking/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/AntoineAugusti/modulus-checking)
[![Coverage Status](http://codecov.io/github/AntoineAugusti/modulus-checking/coverage.svg?branch=master)](http://codecov.io/github/AntoineAugusti/modulus-checking?branch=master)

## Modulus checking
Modulus checking is a procedure for validating sort code and account number combinations. It doesn't confirm that an account belongs to a customer or supports Direct Debit.

If you want to know more about modulus checking, read this [GoCarless guide](https://gocardless.com/guides/posts/modulus-checking/).

## Validity
This package follows the Vocalink specification, version 3.50, that will be live on 7/12/2015. More information about the specification can be seen on the Vocalink website: https://www.vocalink.com/customer-support/modulus-checking

## Included data files
This package ships with the latest version of the modulus weight table data and the sorting code substitution data. Both files can be found in the `data` folder.

## Getting started
You can grab this package with the following command:
```
go get github.com/antoineaugusti/modulus-checking
```

## Usage
If you wanna use the default file parser:
```go
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
```
