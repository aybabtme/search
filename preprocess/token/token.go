package token

import (
	"github.com/aybabtme/search/term"
	"io"
)

// Tokenizer breaks a stream of runes into a bag of terms.
type Tokenizer interface {
	Tokenize(io.Reader) (term.Bag, error)
}
