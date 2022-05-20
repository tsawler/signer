[![Go Report Card](https://goreportcard.com/badge/github.com/tsawler/signer)](https://goreportcard.com/report/github.com/tsawler/signer)
[![Version](https://img.shields.io/badge/goversion-1.18.x-blue.svg)](https://golang.org)
<a href="https://golang.org"><img src="https://img.shields.io/badge/powered_by-Go-3362c2.svg?style=flat-square" alt="Built with GoLang"></a>
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/tsawler/goblender/master/LICENSE)
![Tests](https://github.com/tsawler/signer/actions/workflows/tests.yml/badge.svg)
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="104" height="20" role="img" aria-label="coverage: 100%"><title>coverage: 100%</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="104" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="61" height="20" fill="#555"/><rect x="61" width="43" height="20" fill="#97ca00"/><rect width="104" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110"><text aria-hidden="true" x="315" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="510">coverage</text><text x="315" y="140" transform="scale(.1)" fill="#fff" textLength="510">coverage</text><text aria-hidden="true" x="815" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="330">100%</text><text x="815" y="140" transform="scale(.1)" fill="#fff" textLength="330">100%</text></g></svg>

# Signer

Signer is a simple package that makes signing URLs painless. It uses
[github.com/bwmarrin/go-alone](https://github.com/bwmarrin/go-alone) to sign URLs.

This is useful for things like sending an email with a link that can be verified, and is
tamper-proof.

## Installation

`go get github.com/tsawler/signer@latest`

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
	sign, _ := signer.Signature{Secret: secret}

	// Call the SignURL to get a signed version. Note that only the part after https:// 
	// or http:// is actually signed, but you must pass the full url. This way, we 
	// can use the package in development without worrying about the domain name of 
	// a particular site.
	signed, _ := sign.SignURL("http://example.com/test?id=1")
	fmt.Println("Signed url:", signed)
	
	// output is http://example.com/test?id=1&hash=.3w4TgJ.pAJWBPAO5k1cimZJ-nrRKnlvosOY1Krrp3ALf1rOAds
	
	// verify that a signed URL is valid, and was  issued by this application. Here, 
	// valid is true if the URL has a valid signature, and false if it is not.
	valid, _ := sign.VerifyURL(signed)
	fmt.Println("Valid url:", valid)

	// you can also check for expiry. Here, the signed url expires after 30 minutes.
	expired := sign.Expired(signed, 30)
	fmt.Println("Expired:", expired)
}
```