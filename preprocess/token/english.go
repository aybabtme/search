package token

import (
	"bufio"
	"bytes"
	"github.com/aybabtme/search/term"
	"io"
	"unicode"
)

// English is a tokenizer for the english language.
type English struct {
	TermBagFactory func() term.Bag
}

func (e *English) init() {
	if e.TermBagFactory == nil {
		e.TermBagFactory = DefaultTermBagFactory
	}
}

// Tokenize breaks the content of `r` into terms.
func (e *English) Tokenize(r io.Reader) (term.Bag, error) {
	e.init()
	terms := e.TermBagFactory()
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fields := bytes.FieldsFunc(scanner.Bytes(), func(r rune) bool {
			if r == '-' {
				return false
			}
			if unicode.IsSymbol(r) {
				return true
			}
			if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
				return true
			}
			return false
		})
		for _, field := range fields {
			terms.Add(term.T(field))
		}

	}

	return terms, scanner.Err()
}
