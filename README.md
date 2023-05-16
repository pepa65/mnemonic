# Mnemonic v1.1.0
[![Go Report Card](https://goreportcard.com/badge/github.com/pepa65/mnemonic)](https://goreportcard.com/report/github.com/pepa65/mnemonic)
[![GoDoc](https://godoc.org/github.com/pepa65/mnemonic?status.svg)](https://godoc.org/github.com/pepa65/mnemonic)

**A BIP 39 implementation in Go**

### Features
* Generating human readable sentences for seed generation after [BIP 32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
* All languages mentioned in the [proposal](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) supported:
  English, Japanese, Korean, Spanish, ChineseSimplified, ChineseTraditional, French, Italian
* 128/160/192/224/256 bit (12/15/18/21/24 words) entropy.

## [`mnemonic`](https://godoc.org/github.com/pepa65/mnemonic) package
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

	M "github.com/pepa65/mnemonic"
)

func main() {
	// Generate a random mnemonic in English with 256 bits of entropy (24 words)
	m, _ := M.NewRandom(256, mnemonic.English)

	// Print the mnemonic as a sentence
	fmt.Println(m.Sentence())
	// Print the underlying words
	fmt.Println(m.Words)

	// Validate the mnemonic
	valid, _ := M.IsMnemonicValid(M.English, m.Sentence())
	fmt.Println(valid)

	// Generate a hex-encoded seed from the mnemonic based on "passphrase"
	seed := m.GenerateSeed("passphrase")
	// Print the seed
	fmt.Println(seed)
}
```

## [`entropy`](https://godoc.org/github.com/pepa65/mnemonic/entropy) package
* Supports generating random entropy in the range of 128-256 bits
* Supports generating entropy from a hex string

### Example
```go
package main

import (
	"fmt"

	M "github.com/pepa65/mnemonic"
	E "github.com/pepa65/mnemonic/entropy"
)

func main() {
	// Generate some entropy from a hex string
	ent, _ := E.FromHex("8197a4a47f0425faeaa69deebc05ca29c0a5b5cc76ceacc0")

	// Generate a mnemonic in Japanese with the generated entropy
	jp, _ := M.New(ent, M.Japanese)

	// Print the mnemonic as a sentence
	fmt.Println(jp.Sentence())

	// Generate some random 256 bit entropy
	rnd, _ := E.Random(256)

	// Generate a mnemonic in Spanish with the generated entropy
	sp, _ := M.New(rnd, M.Spanish)

	// Print the mnemonic as a sentence
	fmt.Println(sp.Sentence())
}
```

# Installation
To install `mnemonic` to use as a module, use `go get`:

`go get -u github.com/pepa65/mnemonic`

This will then make the following packages available to you:
* `github.com/pepa65/mnemonic`
* `github.com/pepa65/mnemonic/entropy`

# Contributing
Please feel free to submit issues, fork the repository and send pull requests!
