package stopword

import (
	"github.com/aybabtme/search/word"
)

// Checker is something that can tell if a word is a stopword.
type Checker interface {
	IsStopWord(word word.W) bool
}
