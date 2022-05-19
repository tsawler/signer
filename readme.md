# Signer

Signer is a simple package that makes signing URLs painless. It uses
[github.com/bwmarrin/go-alone](https://github.com/bwmarrin/go-alone) to sign URLs.

## Installation

`go get -u github.com/tsawler/signer`

## Usage

```golang

package main

import (
	"fmt"
	"github.com/tsawler/signer"
)

const secret = "somelongsecuresecret"

func main() {
	// create a variable of type Signature, and pass it a secret (<= 64 characters)
	sign := signer.Signature{Secret: secret}

	// call the SignURL to get a signed version. Note that only the part after https:// or http:// is actually signed,
	// but you must pass the full url. This way, we can use the package in development without worrying about the 
	// domain name of a particular site.
	signed := sign.SignURL("http://example.com/test?id=1")
	fmt.Println("Signed url:", signed)

	// verify that a signed URL is valid, and was  issued by this application. Here, valid is true if the URL has a 
	// valid signature, and false if it is not.
	valid := sign.VerifyURL(signed)
	fmt.Println("Valid url:", valid)

	// you can also check for expiry. Here, the signed url expires after 30 minutes.
	expired := sign.Expired(signed, 30)
	fmt.Println("Expired:", expired)
}
```