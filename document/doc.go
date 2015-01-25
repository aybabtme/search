package document

import (
	"github.com/aybabtme/search/word"
)

// Doc is a document.
type Doc interface {
	Keywords() word.Set
}
