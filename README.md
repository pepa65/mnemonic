# Mnemonic v1.0.1
[![Build Status](https://travis-ci.org/brianium/mnemonic.svg?branch=master)](https://travis-ci.org/brianium/mnemonic)
[![Go Report Card](https://goreportcard.com/badge/github.com/pepa65/mnemonic)](https://goreportcard.com/report/github.com/pepa65/mnemonic)
[![GoDoc](https://godoc.org/github.com/brianium/mnemonic?status.svg)](https://godoc.org/github.com/brianium/mnemonic)

**A BIP 39 implementation in Go**

### Features
* Generating human readable sentences for seed generation after [BIP 32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
* All languages mentioned in the [proposal](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) supported.
* 128/160/192/224/256 bit (12/15/18/21/24 words) entropy.

## [`mnemonic`](https://godoc.org/github.com/brianium/mnemonic) package
* Generates human readable sentences and the seeds derived from them.
* Supports all languages mentioned in the [BIP 39 proposal](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki).
* Supports ideographic spaces for Japanese language.
* Forked from github.com/brianium/mnemonic
* Generates the right amount of words

### Example

```go
package main

import (
	"fmt"

	"github.com/pepa65/mnemonic"
)

func main() {
	// generate a random Mnemonic in English with 256 bits of entropy (24 words)
	m, _ := mnemonic.NewRandom(256, mnemonic.English)

	// print the Mnemonic as a sentence
	fmt.Println(m.Sentence())

	// validate Mnemonic
	valid, _ := mnemonic.IsMnemonicValid(mnemonic.English, m.Sentence())
	fmt.Println(valid)

	// inspect underlying words
	fmt.Println(m.Words)

	// generate a seed from the Mnemonic
	seed := m.GenerateSeed("passphrase")

	// print the seed as a hex encoded string
	fmt.Println(seed)
}
```

## [`entropy`](https://godoc.org/github.com/brianium/mnemonic/entropy) package
* Supports generating random entropy in the range of 128-256 bits
* Supports generating entropy from a hex string

### Example

```go
package main

import (
	"fmt"

	"github.com/pepa65/mnemonic"
	"github.com/pepa65/mnemonic/entropy"
)

func main() {
	// generate some entropy from a hex string
	ent, _ := entropy.FromHex("8197a4a47f0425faeaa69deebc05ca29c0a5b5cc76ceacc0")

	// generate a Mnemonic in Japanese with the generated entropy
	jp, _ := mnemonic.New(ent, mnemonic.Japanese)

	// print the Mnemonic as a sentence
	fmt.Println(jp.Sentence())

	// generate some random 256 bit entropy
	rnd, _ := entropy.Random(256)

	// generate a Mnemonic in Spanish with the generated entropy
	sp, _ := mnemonic.New(rnd, mnemonic.Spanish)

	// print the Mnemonic as a sentence
	fmt.Println(sp.Sentence())
}
```

# Installation

To install Mnemonic, use `go get`:

`go get -u github.com/pepa65/mnemonic`

This will then make the following packages available to you:
* `github.com/pepa65/mnemonic`
* `github.com/pepa65/mnemonic/entropy`

Import the `mnemonic` package into your code using this template:

```go
package yours

import (
	"github.com/pepa65/mnemonic"
)

func MnemonicJam(passphrase string) {
	m := mnemonic.NewRandom(passphrase)
}
```

# Contributing

Please feel free to submit issues, fork the repository and send pull requests!
When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.
