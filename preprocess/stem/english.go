package stem

import (
	"github.com/dchest/stemmer/porter2"

	"github.com/aybabtme/search/term"
)

// English stems words of the english language.
type English struct {
}

// Stem a term to get its root.
func (e English) Stem(t term.T) term.T {
	root := porter2.Stemmer.Stem(string(t))
	return term.T(root)
}
