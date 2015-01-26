package token

import (
	"github.com/aybabtme/search/term"
	"io"
)

// DefaultTermBagFactory create using a term.HashBag implementation.
func DefaultTermBagFactory() term.Bag { return new(term.HashBag) }

// Tokenizer breaks a stream of runes into a bag of terms.
type Tokenizer interface {
	Tokenize(io.Reader) (term.Bag, error)
}
