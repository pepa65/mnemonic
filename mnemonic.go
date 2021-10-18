package mnemonic

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pepa65/mnemonic/entropy"
)

// Mnemonic represents a collection of human readable words
// used for HD wallet seed generation
type Mnemonic struct {
	Words    []string
	Language Language
}

// New returns a new Mnemonic for the given entropy and language
func New(ent []byte, lang Language) (*Mnemonic, error) {
	const chunkSize = 11
	bits := entropy.CheckSummed(ent)
	length := len(bits)
	words := make([]string, length/11)
	for i := 0; i < length; i += chunkSize {
		stringVal := string(bits[i : chunkSize+i])
		intVal, err := strconv.ParseInt(stringVal, 2, 64)
		if err != nil {
			return nil, fmt.Errorf("Could not convert %s to word index", stringVal)
		}
		word, err := GetWord(lang, intVal)
		if err != nil {
			return nil, err
		}
		words[(chunkSize+i)/11-1] = word
	}
	m := Mnemonic{words, lang}
	return &m, nil
}

// NewRandom returns a new Mnemonic with random entropy of the given length
// in bits
func NewRandom(length int, lang Language) (*Mnemonic, error) {
	ent, err := entropy.Random(length)
	if err != nil {
		return nil, fmt.Errorf("Error generating random entropy: %s", err)
	}
	return New(ent, lang)
}

func IsMnemonicValid(lang Language, sentence string) (bool, error) {
	words := strings.Split(sentence, " ")
	buff := bytes.NewBuffer(nil)
	for _, word := range words {
		idx, err := GetIndex(lang, word)
		if err != nil {
			fmt.Println(word)
			//return false, err
		}
		buff.Write(paddingLeft([]byte(strconv.FormatInt(idx, 2))))
	}
	ent, chksum := entropy.BitsToBytes(buff.Bytes())
	if ent == nil && chksum == nil {
		return false, entropy.ErrBitsChecksumLength
	}
	return bytes.Equal(entropy.CheckSum(ent), chksum), nil
}

func paddingLeft(data []byte) []byte {
	const chunkSize = 11
	ret := make([]byte, chunkSize)
	length := len(data)
	if length < chunkSize {
		copy(ret[chunkSize-length:], data)
		return ret[:]
	} else {
		return data[:chunkSize]
	}
}

// Sentence returns Mnemonic's words as a space separated sentence
func (m *Mnemonic) Sentence() string {
	if m.Language == Japanese {
		return strings.Join(m.Words, `　`)
	}
	return strings.Join(m.Words, " ")
}

func RecoverFromMnemonic(lang Language, sentence string) (m *Mnemonic, err error) {
	var check bool
	if check, err = IsMnemonicValid(lang, sentence); err != nil {
		return nil, err
	}
	if !check {
		return nil, errors.New("invalid mnemonic")
	}
	words := strings.Split(sentence, " ")
	m = &Mnemonic{
		Words:    words,
		Language: lang,
	}
	return m, nil
}

// GenerateSeed returns a seed used for wallet generation per
// BIP-0032 or similar method. The internal Words set
// of the Mnemonic will be used
func (m *Mnemonic) GenerateSeed(passphrase string) *Seed {
	return NewSeed(m.Sentence(), passphrase)
}
