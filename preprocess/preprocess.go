package preprocess

import (
	"github.com/aybabtme/search/term"
	"io"
)

// Preprocessor cleans up the content of a reader and
// returns a bag of terms.
type Preprocessor interface {
	Process(io.Reader) (term.Bag, error)
}
