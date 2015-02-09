package preprocess

import (
	"github.com/aybabtme/search/preprocess/stem"
	"github.com/aybabtme/search/preprocess/stopword"
	"github.com/aybabtme/search/preprocess/token"
	"github.com/aybabtme/search/term"
	"io"
)

// English is a processor for the english language. By default,
// it uses a term.DefaultBagFactory term bag factory, and an
// english stemmer, stop word checker and tokenizer.
type English struct {
	TermBagFactory func() term.Bag
	Stemmer        *stem.English
	Checker        *stopword.English
	Tokenizer      *token.English
}

func (e *English) init() {
	if e.TermBagFactory == nil {
		e.TermBagFactory = term.DefaultBagFactory
	}
	if e.Stemmer == nil {
		e.Stemmer = new(stem.English)
	}
	if e.Checker == nil {
		e.Checker = new(stopword.English)
	}
	if e.Tokenizer == nil {
		e.Tokenizer = new(token.English)
	}
}

// Process the content of r and returns a bag of term
// representing it.
func (e *English) Process(r io.Reader) (term.Bag, error) {
	e.init()
	tokens, err := e.Tokenizer.Tokenize(r)
	processed := e.TermBagFactory()
	tokens.Iter(func(t term.T) {
		if e.Checker.IsStopWord(t) {
			return // ignore
		}
		processed.Add(e.Stemmer.Stem(t)) // but insert the stem
	})
	return processed, err
}
