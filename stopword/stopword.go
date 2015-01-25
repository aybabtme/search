package stopword

import (
	"github.com/aybabtme/search/term"
)

// Checker is something that can tell if a term is a stopword.
type Checker interface {
	IsStopWord(term term.T) bool
}
