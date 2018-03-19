[![Travis CI](https://img.shields.io/travis/AntoineAugusti/moduluschecking/master.svg?style=flat-square)](https://travis-ci.org/AntoineAugusti/moduluschecking)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/AntoineAugusti/moduluschecking/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/AntoineAugusti/moduluschecking)
[![Coverage Status](http://codecov.io/github/AntoineAugusti/moduluschecking/coverage.svg?branch=master)](http://codecov.io/github/AntoineAugusti/moduluschecking?branch=master)

## Modulus checking
Modulus checking is a procedure for validating sort code and account number combinations. It doesn't confirm that an account belongs to a customer or supports Direct Debit.

If you want to know more about modulus checking, read this [GoCardless guide](https://gocardless.com/guides/posts/modulus-checking/).

## Validity
This package follows the Vocalink specification, version 4.80, that will be live on 16/04/2018. More information about the specification can be seen on the [Vocalink website](https://www.vocalink.com/customer-support/modulus-checking).

## API
If you prefer to send request to a web service, take a look at the package [moduluschecking-api](https://github.com/AntoineAugusti/moduluschecking-api) that offers an API to validate UK bank account numbers, supporting authentication and rate limits.

## Institution covered
The following institutions are supported:

- Allied Irish
- Bank of England
- Bank of Ireland
- Bank of Scotland
- Barclays
- Bradford and Bingley Building Society
- Charity Bank
- Citibank
- ClearBank
- Clydesdale
- Co-Operative Bank
- Contis Financial Services
- Coutts
- First Trust
- HSBC
- HSBC
- Halifax
- Hoares Bank
- Lloyds
- Metro Bank
- NatWest
- Nationwide Building Society
- Northern
- Orwell Union Ltd
- Royal Bank of Scotland
- Santander
- Secure Trust
- Starling Bank
- TSB
- Tesco Bank
- Ulster Bank
- Unity Trust Bank
- Virgin Bank
- Woolwich
- Yorkshire Bank

## Included data files
This package ships with the latest version of the modulus weight table data and the sorting code substitution data. Both files can be found in the `data` folder.

## Getting started
You can grab this package with the following command:
```
go get github.com/AntoineAugusti/moduluschecking/...
```

## Usage
If you wanna use the default file parser:
```go
package main

import (
    "fmt"

    "github.com/AntoineAugusti/moduluschecking/models"
    "github.com/AntoineAugusti/moduluschecking/parsers"
    "github.com/AntoineAugusti/moduluschecking/resolvers"
)

func main() {
    // Read the modulus weight table and the sorting
    // code substitution table from the folder data
    parser := parsers.CreateFileParser()

    // The resolver handles the verification of the validity of
    // bank accounts according to the data obtained by the parser
    resolver := resolvers.NewResolver(parser)

    // This helper method handles special cases for
    // bank accounts from:
    // - National Westminster Bank plc (10 or 11 digits with possible presence of dashes, for account numbers)
    // - Co-Operative Bank plc (10 digits for account numbers)
    // - Santander (9 digits for account numbers)
    // - banks with 6 or 7 digits for account numbers
    bankAccount := models.CreateBankAccount("089999", "66374958")

    // Check if the created bank account is valid against the rules
    fmt.Println(resolver.IsValid(bankAccount))
}
```

## Benchmark
On my personal laptop (MacBook Pro, Core i5 2.5 Ghz, 8 GB of RAM with a SSD):
- reading data files from the filesytem by creating the parser and the resolver: ~350 ms
- checking the validity of 1,000 bank account numbers: ~7 ms
