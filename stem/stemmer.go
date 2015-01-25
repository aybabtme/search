package stem

import "github.com/aybabtme/search/term"

// Stemmer can take terms and return their root.
type Stemmer interface {
	Stem(term term.T) (root term.T)
}
